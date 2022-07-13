.PHONY: set-perf-event-paranoid
set-perf-event-paranoid:
	echo 1 | sudo tee /proc/sys/kernel/perf_event_paranoid

.PHONY: check-perf-list
check-perf-list:
	perf list