package genKill

type GenKillTask interface{
	InitMaps()
	InitGenKillMap()
	InitBeforeAfterMap()
	Analyze()
}
