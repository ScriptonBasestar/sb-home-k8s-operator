package helper

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"log"
	"os"
)

func PullHelmChart(repoUrl string, repoName string, chartName string, chartVersion string) error {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), "default", "secret", log.Printf); err != nil {
		return err
	}

	client := action.NewPull()
	client.RepoURL = repoUrl
	//client.RepoName = repoName
	//client.ChartName = chartName
	client.Version = chartVersion
	client.DestDir = os.Getenv("HELM_DIR")

	_, err := client.Run(chartName)
	return err
}

//func LintHelmChart(destDir string, chartName string, yamlString map[string]interface{}) error {
//	client := action.NewLint()
//	client.Run(destDir, yamlString)
//	return nil
//}

//func Template() error {
//	return nil
//}

func InstallHelmChart(repoUrl string, repoName string, chartName string) error {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), "default", "secret", log.Printf); err != nil {
		return err
	}

	client := action.NewInstall(actionConfig)
	chartPath := "/path/to/chart"
	chart, err := loader.Load(chartPath)
	if err != nil {
		return err
	}

	_, err = client.Run(chart, nil)
	return err
}

func UninstallHelmChart() error {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), "default", "secret", log.Printf); err != nil {
		return err
	}

	client := action.NewUninstall(actionConfig)

	_, err := client.Run("my-release")
	return err
}

func UpgradeHelmChart() error {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), "default", "secret", log.Printf); err != nil {
		return err
	}

	client := action.NewUpgrade(actionConfig)
	chartPath := "/path/to/chart"
	chart, err := loader.Load(chartPath)
	if err != nil {
		return err
	}

	_, err = client.Run("my-release", chart, nil)
	return err
}
