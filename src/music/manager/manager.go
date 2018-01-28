package manager 
import (
	"errors"
)

type MusicEntry struct{
	Id string 
	Name string
	Artist string 
	Source string 
	Type string
}

type MusicManager struct{
	musics []MusicEntry
}

func NewMusicManager() *MusicManager{  //由于存储音乐文件的信息结构太过庞大，所以这里采用的是指针的方式进行引用
	return &MusicManager{make([]MusicEntry, 0)} //这里创建了一个音乐文件管理结构，这里虽然创建的切片大小为0，但是这里用make给该结构体创建了一个地址，接下来就可以通过改地址来往这个结构中添加数据了
}

func (m *MusicManager) Len() int{
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error){
	if index <0 || index >= len(m.musics){
		return nil, errors.New("Index out of range")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (music *MusicEntry, err error){

	if len(m.musics)==0{
		return nil, errors.New(" no item to Find") //这里使用erros模块来自定义一个错误值用来返回
	}

	for _, m := range m.musics{
		if m.Name == name{
			return &m, nil
		}
	}

	return nil, errors.New(" no item to Find")

}

func (m *MusicManager) Add(music *MusicEntry){
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry{
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removemusic := &m.musics[index]

	m.musics = append(m.musics[:index], m.musics[index+1:] ...)

	return removemusic 
}


