package agent

import (
)

type SpatialMap struct {
	space map[string]*Zone
}

type Zone struct {
	ents []Entity
}

func CreateSpatialMap() SpatialMap{
	return SpatialMap{make(map[string]*Zone)}
}

func CreateZone() *Zone {
	return &Zone{make([]Entity, 0)}
}

func CreateZoneOne(ent Entity) *Zone {
	zone := Zone{make([]Entity, 1)}
	zone.ents[0] = ent
	return &zone
}

func (spatial *SpatialMap) SpatialGetZone(ent Entity) []Entity {
	key := getPosKey(ent)
	if val, ok := spatial.space[key]; ok {
		return val.ents
	} else {
		return nil
	}
}

func (spatial *SpatialMap) SpatialAdd(ent Entity) {
	key := getPosKey(ent)

	if val, ok := spatial.space[key]; ok {
		val.ZoneAdd(ent)
	} else {
		spatial.space[key] = CreateZoneOne(ent)
	}
}

func (spatial *SpatialMap) SpatialReset() {
	spatial.space = make(map[string]*Zone)
}

func (zone *Zone) ZoneAdd(new_ent Entity) {
	zone.ents = append(zone.ents, new_ent)
}

func getPosKey(ent Entity) string {
	pos := ent.GetPos()
	res := 1000

	x_key := string((int(pos.X) / res) * res)
	y_key := string((int(pos.Y) / res) * res)
	z_key := string((int(pos.Z) / res) * res)

	return x_key + "-" + y_key + "-" + z_key
}












