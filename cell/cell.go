package cell

type Cell struct {
	isAlive  bool;
}

func NewCell(isAlive bool) Cell {
 
	return Cell{isAlive: isAlive}

}

func (c Cell) IsAlive() bool {
	return c.isAlive;
}

func (c Cell) NextGenerationCell(noOfLiveNeighbours int) Cell{

    if(c.IsAlive()){
         if(noOfLiveNeighbours >= 2 && noOfLiveNeighbours <= 3){
			return NewCell(true);
		 }
	}
	if(noOfLiveNeighbours == 3){
		return NewCell(true)
	}
	return NewCell(false);

}
