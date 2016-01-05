package core

var (
	ActionNone                    string = ""
	ActionContainerArchive        string = "container_archive"
	ActionContainerArchiveHead    string = "container_archive_head"
	ActionContainerArchiveExtract string = "container_archive_extract"
	ActionContainerAttach         string = "container_attach"
	ActionContainerAttachWs       string = "container_attachws"
	ActionContainerCreate         string = "container_create"
	ActionContainerCommit         string = "container_commit"
	ActionContainerDelete         string = "container_delete"
	ActionContainerExecCreate     string = "container_exec_create"
	ActionContainerExecStart      string = "container_exec_start"
	ActionContainerExecInspect    string = "container_exec_inspect"
	ActionContainerExport         string = "container_export"
	ActionContainerCopyFiles      string = "container_copyfiles"
	ActionContainerStart          string = "container_start"
	ActionContainerRestart        string = "container_restart"
	ActionContainerRename         string = "container_rename"
	ActionContainerKill           string = "container_kill"
	ActionContainerStop           string = "container_stop"
	ActionContainerPause          string = "container_pause"
	ActionContainerUnpause        string = "container_unpause"
	ActionContainerInspect        string = "container_inspect"
	ActionContainersList          string = "container_list"
	ActionContainerLogs           string = "container_logs"
	ActionContainerTop            string = "container_top"
	ActionContainerChanges        string = "container_changes"
	ActionContainerStats          string = "container_stats"
	ActionContainerResize         string = "container_resize"
	ActionContainerWait           string = "container_wait"
	ActionImageBuild              string = "image_build"
	ActionImageCreate             string = "image_create"
	ActionImageDelete             string = "image_delete"
	ActionImagePush               string = "image_push"
	ActionImagesLoad              string = "images_load"
	ActionImageInspect            string = "image_inspect"
	ActionImagesList              string = "image_list"
	ActionImageHistory            string = "image_history"
	ActionImageTag                string = "image_tag"
	ActionImagesSearch            string = "images_search"
	ActionImagesArchive           string = "images_archive"
	ActionDockerEvents            string = "events"
	ActionDockerVersion           string = "version"
	ActionDockerCheckAuth         string = "auth"
	ActionDockerInfo              string = "info"
	ActionDockerPing              string = "ping"
)
