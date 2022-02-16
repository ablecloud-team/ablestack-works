package main

func destroyWorkspaces(workspaceInfo Workspace) {
	ResultDeleteGroup, errDeleteGroup := deleteGroup(workspaceInfo.Name)
	ResultDeleteWorkspace := deleteWorkspace(workspaceInfo.Uuid)
	log.Infof("%v%v%v", ResultDeleteGroup, errDeleteGroup, ResultDeleteWorkspace)
}
