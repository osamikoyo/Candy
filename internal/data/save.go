package data

func (p *Post) Save(d *Database) error {
	return d.Create(p).Error
}
