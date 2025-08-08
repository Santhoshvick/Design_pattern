package main

type Course struct{
	CourseName string
	CourseDuration string
	FacultyName string
	Timing string
	Videos string
	quiz string
	Assignment string
	Materails string
	Certificate string
}

type courseBuilder interface{
	addVideos(string,string)*courseBuilder
	addQuizzes(string)*courseBuilder
	addAssignments(string)*courseBuilder
	addReadingMaterial(string)*courseBuilder
	createCertificate(string)*courseBuilder
	build()Course
}


type JavaFullStack struct{
	course Course
}

func(j *JavaFullStack)addVideos(name string,videos string)*JavaFullStack{
	j.course.CourseName=name
	j.course.Videos=videos
	return j
}

func(j *JavaFullStack)addQuizzes(quizContent string)*JavaFullStack{
	j.course.quiz=quizContent
	return j
}

func(j *JavaFullStack)addAssignments(assignment string)*JavaFullStack{
	j.course.Assignment=assignment
	return j
}

func(j *JavaFullStack)addReadingMaterial(material string)*JavaFullStack{
	j.course.Materails=material
	return j
}

func(j *JavaFullStack)createCertificate(certificate string)*JavaFullStack{
	j.course.Certificate=certificate
	return j
}

func(j *JavaFullStack)build()Course{
	return j.course
}





type PythonFullStack struct{
	course Course
}

func(p *PythonFullStack)addVideos(name string,videos string)*PythonFullStack{
	p.course.CourseName=name
	p.course.Videos=videos
	return p
}

func(p *PythonFullStack)addQuizzes(quizContent string)*PythonFullStack{
	p.course.quiz=quizContent
	return p
}

func(p *PythonFullStack)addAssignments(assignment string)*PythonFullStack{
	p.course.Assignment=assignment
	return p
}

func(p *PythonFullStack)addReadingMaterial(material string)*PythonFullStack{
	p.course.Materails=material
	return p
}

func(p *PythonFullStack)createCertificate(certificate string)*PythonFullStack{
	p.course.Certificate=certificate
	return p
}

func(p *PythonFullStack)build()Course{
	return p.course
}