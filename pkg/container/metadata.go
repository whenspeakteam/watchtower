package container

const (
	watchtowerLabel = "com.centurylinklabs.watchtower"
	signalLabel     = "com.centurylinklabs.watchtower.stop-signal"
	enableLabel     = "com.centurylinklabs.watchtower.enable"
	zodiacLabel     = "com.centurylinklabs.zodiac.original-image"
	preCheckLabel   = "com.centurylinklabs.watchtower.lifecycle.pre-check"
	postCheckLabel  = "com.centurylinklabs.watchtower.lifecycle.post-check"
	preUpdateLabel  = "com.centurylinklabs.watchtower.lifecycle.pre-update"
	postUpdateLabel = "com.centurylinklabs.watchtower.lifecycle.post-update"

	updaterContainerLabel = "com.centurylinklabs.watchtower.lifecycle.updater"
	workDirLabel = "com.centurylinklabs.watchtower.lifecycle.workdir"
	servicesLabel = "com.centurylinklabs.watchtower.lifecycle.services"
)

// GetLifecyclePreCheckCommand returns the pre-check command set in the container metadata or an empty string
func (c Container) GetLifecyclePreCheckCommand() string {
	return c.getLabelValueOrEmpty(preCheckLabel)
}

// GetLifecyclePostCheckCommand returns the post-check command set in the container metadata or an empty string
func (c Container) GetLifecyclePostCheckCommand() string {
	return c.getLabelValueOrEmpty(postCheckLabel)
}

// GetLifecyclePreUpdateCommand returns the pre-update command set in the container metadata or an empty string
func (c Container) GetLifecyclePreUpdateCommand() string {
	return c.getLabelValueOrEmpty(preUpdateLabel)
}

// GetLifecyclePostUpdateCommand returns the post-update command set in the container metadata or an empty string
func (c Container) GetLifecyclePostUpdateCommand() string {
	return c.getLabelValueOrEmpty(postUpdateLabel)
}

func (c Container) GetLifecycleUpdaterContainer() string {
	return c.getLabelValueOrEmpty(updaterContainerLabel)
}

func (c Container) GetLifecycleWorkDir() string {
	return c.getLabelValueOrEmpty(workDirLabel)
}

func (c Container) GetLifecycleServices() string {
	return c.getLabelValueOrEmpty(servicesLabel)
}

// ContainsWatchtowerLabel takes a map of labels and values and tells
// the consumer whether it contains a valid watchtower instance label
func ContainsWatchtowerLabel(labels map[string]string) bool {
	val, ok := labels[watchtowerLabel]
	return ok && val == "true"
}

func (c Container) getLabelValueOrEmpty(label string) string {
	if val, ok := c.containerInfo.Config.Labels[label]; ok {
		return val
	}
	return ""
}

func (c Container) getLabelValue(label string) (string, bool) {
	val, ok := c.containerInfo.Config.Labels[label]
	return val, ok
}
