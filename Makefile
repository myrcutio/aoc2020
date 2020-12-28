.PHONY: csvTemplate
csvTemplate:
	boilr template save ./boilr/csvData aocday -f

.PHONY: multilineTemplate
multilineTemplate:
	boilr template save ./boilr/multilineData aocday -f

pkg/day%:
	boilr template use aocday pkg <<< $*
