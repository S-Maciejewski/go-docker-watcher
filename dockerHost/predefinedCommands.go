package dockerHost

var licensePaths = map[string]string{
	"kibana":        "/usr/share/kibana/LICENSE.txt",
	"elasticsearch": "/usr/share/elasticsearch/LICENSE.txt",
}

func GetLsForContainer(containerName string) string {
	return GetCustomCommandResult(containerName, []string{"ls", "-l"})
}

func GetReadmeForContainer(containerName string) string {
	return GetCustomCommandResult(containerName, []string{"cat", "/README.txt"})
}

func GetLicenseForContainer(containerName string) string {
	path := licensePaths[containerName]
	if path == "" {
		return "No license path found for container: " + containerName
	}
	return GetCustomCommandResult(containerName, []string{"cat", path})
}
