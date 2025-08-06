package main

type SmartHouseSystem struct{
	Tv *tv
	Light *light
	MusicSystem *musicSytem
}


func newHouseSystem(tv1 *tv,light1 *light,musicsystem1 *musicSytem)*SmartHouseSystem{
	return &SmartHouseSystem{
		Tv: tv1,
		Light: light1,
		MusicSystem: musicsystem1,
	}
}

func (s *SmartHouseSystem)Turnon(){

	s.Tv.tvOn()
	s.Light.lightsOn()
	s.MusicSystem.musicOn()
}

func (s *SmartHouseSystem)Turnoff(){
	s.Tv.tvOff()
	s.Light.lightsOff()
	s.MusicSystem.musicOff()
}