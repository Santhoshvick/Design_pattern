package main



type PcItem struct{
	Monitor string
	Keyboard string
	Mouse string
	CPU string
	RAM string
	GPU string
	Storage string
}

type PCBuilder interface{
	addMonitor()
	addKeyboard()
	addCpu()
	addMouse()
}

type NormalPc struct{
	pcitem PcItem
}

func(n *NormalPc)addMonitor(monitor string)*NormalPc{
   n.pcitem.Monitor=monitor
   return n
}

func(n *NormalPc)addKeyboard(keyboard string)*NormalPc{
	n.pcitem.Keyboard=keyboard
	return n
}

func(n *NormalPc)addMouse(mouse string)*NormalPc{
	n.pcitem.Mouse=mouse
	return  n
}

func(n *NormalPc)addCpu(storage string,ram string,gpu string)*NormalPc{
	n.pcitem.Storage=storage
	n.pcitem.RAM=ram
	n.pcitem.GPU=gpu
	return n
}

