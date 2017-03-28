package reading

import gin "gopkg.in/gin-gonic/gin.v1"

func (p *Engine) getStatus(c *gin.Context) (interface{}, error) {

	var bc int
	if err := p.Db.Model(&Book{}).Count(&bc).Error; err != nil {
		return nil, err
	}

	dict := gin.H{}
	for _, dic := range dictionaries {
		dict[dic.GetBookName()] = dic.GetWordCount()
	}

	return gin.H{
		"books": gin.H{"count": bc},
		"dict":  dict,
	}, nil
}
