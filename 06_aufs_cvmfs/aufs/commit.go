package aufs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

func (a *Driver) getParentThinLayer(id string) (ThinImage, error) {
	roLayers, _ := getParentIDs(a.rootPath(), id)
	var thin ThinImage

	for _, l := range roLayers {
		if a.isThinImageLayer(l) {
			return a.readThinFile(l), nil
		}
	}

	return thin, errors.New("Not a thin image!")
}

func (a *Driver) readThinFile(id string) ThinImage {
	thin_file_path := path.Join(a.getDiffPath(id), ".thin")
	content, _ := ioutil.ReadFile(thin_file_path)

	var thin ThinImage
	json.Unmarshal(content, &thin)

	return thin
}

func (a *Driver) writeThinFile(thin ThinImage, id string) {
	p := path.Join(a.getDiffPath(id), ".thin")
	j, _ := json.Marshal(thin)
	ioutil.WriteFile(p, j, os.ModePerm)
}
