def command_switch_to_g_running(fn):
	foundg = None
	gs = goroutines().Goroutines
	for g in gs:
		locs = stacktrace(g.ID, 20, False, False, None, None).Locations
		for loc in locs:
			if foundg != None:
				break

			if loc.Function != None and str(loc.Function.Name_) == fn:
				foundg = g
				break

	if foundg != None:
		dlv_command("goroutine %d" % foundg.ID)
	else:
		print("could not find G")