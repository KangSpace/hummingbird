package thread

//var threads map[int]Thread;

func init() {
	//线程集合
	//	threads = make(map[int]Thread)
}

type threadsConstant struct {
	threads map[int]Thread
}

/*
//获取当前线程集合
func getThreads() (map[int]Thread){
	return threads
}

func addThread(id int , t *Thread) (*Thread){
	threads[id] = t
	return t;
}*/
