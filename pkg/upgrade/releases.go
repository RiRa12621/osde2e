package upgrade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Masterminds/semver"
	"github.com/openshift/osde2e/pkg/config"
)

const (
	// format string for release stream latest from release controller
	latestReleaseControllerURLFmt = "https://openshift-release.svc.ci.openshift.org/api/v1/releasestream/%s/latest"
	// format string for Cincinnati releases
	cincinnatiURLFmt = "https://api%s.openshift.com/api/upgrades_info/v1/graph?channel=%s"
)

// LatestRelease retrieves latest release information for given releaseStream. Will use Cincinnati for stage/prod.
func LatestRelease(cfg *config.Config) (name, pullSpec string, err error) {
	releaseStream := cfg.UpgradeReleaseStream

	var resp *http.Response
	var data []byte
	if cfg.OSDEnv == "int" {
		latestURL := fmt.Sprintf(latestReleaseControllerURLFmt, releaseStream)
		resp, err = http.Get(latestURL)
		if err != nil {
			err = fmt.Errorf("failed to get latest for stream '%s': %v", releaseStream, err)
			return
		}

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			err = fmt.Errorf("failed reading body: %v", err)
			return
		}

		latest := latestAccepted{}
		if err = json.Unmarshal(data, &latest); err != nil {
			err = fmt.Errorf("error decoding body of '%s': %v", data, err)
		}

		return latest.Name, latest.PullSpec, nil
	}

	// If stage or prod, use Cincinnati instead of the release controller
	stage := ""

	// Add in stage to the URL if necessary
	if cfg.OSDEnv == "stage" {
		stage = ".stage"
	}

	cincinnatiFormattedURL := fmt.Sprintf(cincinnatiURLFmt, stage, releaseStream)

	var req *http.Request

	// Cincinnati requires an Accept header, so we add it in here
	req, err = http.NewRequest("GET", cincinnatiFormattedURL, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		err = fmt.Errorf("failed to create Cincinnati request for URL '%s': %v", cincinnatiFormattedURL, err)
		return
	}

	resp, err = (&http.Client{}).Do(req)

	if err != nil {
		err = fmt.Errorf("Request failed for URL '%s': %v", cincinnatiFormattedURL, err)
		return
	}

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		err = fmt.Errorf("Failed reading body: %v", err)
		return
	}

	var cincinnatiReleases cincinnatiReleaseNodes
	var latestVersion *semver.Version
	var latestCincinnatiRelease cincinnatiRelease

	if err = json.Unmarshal(data, &cincinnatiReleases); err != nil {
		err = fmt.Errorf("error decoding body of '%s': %v", data, err)
	}

	for _, release := range cincinnatiReleases.Nodes {
		currentVersion, err := semver.NewVersion(release.Version)

		if err != nil {
			log.Printf("Unable to parse version for %s, skipping", release.Version)
			continue
		}

		if latestVersion == nil || currentVersion.GreaterThan(latestVersion) {
			latestVersion = currentVersion
			latestCincinnatiRelease = release
		}
	}

	return latestCincinnatiRelease.Version, latestCincinnatiRelease.Payload, nil
}

// latestAccepted information from release controller.
type latestAccepted struct {
	Name        string `json:"name"`
	PullSpec    string `json:"pullSpec"`
	DownloadURL string `json:"downloadURL"`
}

type cincinnatiReleaseNodes struct {
	Nodes []cincinnatiRelease `json:"nodes"`
}
type cincinnatiRelease struct {
	Version string `json:"version"`
	Payload string `json:"payload"`
}
