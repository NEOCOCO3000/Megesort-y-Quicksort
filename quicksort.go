
paquete quicksort

func  QuickSort ( slice [] int ) {
  r�pido (slice, 0 , len (slice) - 1 )
}

func  quick ( slice [] int , start  int , end  int ) {
  if start <end {
    pivotIndex  : =  divisi�n (corte, inicio, fin)
    r�pido (slice, start, pivotIndex- 1 )
    r�pido (slice, pivotIndex + 1 , final)
  }
}

func  split ( slice [] int , start  int , end  int ) int {
  pivot  : = slice [end]
  j  : = inicio

  para  �ndice  : = inicio; �ndice <fin; �ndice ++ {
    if slice [index] <pivot {
      temp  : = slice [index]
      slice [index] = slice [j]
      slice [j] = temp
      j = j + 1
    }
  }

  si pivot <= slice [j] {
    slice [end] = slice [j]
    rebanada [j] = pivote
  }
  
  devolver j
}