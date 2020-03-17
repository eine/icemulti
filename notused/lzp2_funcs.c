// "LZP2 v1.1 copyright (c) 1995 Charles Bloom\n"

void Encode(void) {
  long MatchLen;
  long Hash,i;
  ubyte *p;
  ubyte *CurRawArrayPtr,*RawArrayPtrDone;

  CurRawArrayPtr = RawArray;
  RawArrayPtrDone = RawArray + RawFileLen;

  NumLiterals = NumLengths = 0;

  for(i=0;i<3;i++)
	Literals[NumLiterals++] = *CurRawArrayPtr++;

  while ( CurRawArrayPtr < RawArrayPtrDone ) {
	Hash = HASHFUNC(CurRawArrayPtr[-1],CurRawArrayPtr[-2],CurRawArrayPtr[-3]);
	p = Table[Hash]; Table[Hash] = CurRawArrayPtr;
	if ( p ) {
	  MatchLen=0;
	  while(*p == *CurRawArrayPtr) { p++; CurRawArrayPtr++; MatchLen++; }
      if ( MatchLen > 0 ) {
        while( MatchLen >= MatchLenEscape ) {
          Lengths[NumLengths++] = MatchLenEscape;
          MatchLen -= MatchLenEscape;
        }
		Lengths[NumLengths++] = (ubyte)MatchLen;
	  }else{ Lengths[NumLengths++] = 0; }
    }
  #ifndef LiteralAfterMatch
	else MatchLen = 0;
	if ( MatchLen == 0 )
  #endif
	Literals[NumLiterals++] = *CurRawArrayPtr++;
  }
}

void Decode(void) {
  long MatchLen;
  long Hash,i;
  ubyte *p;
  ubyte *CurRawArrayPtr,*RawArrayPtrDone;
  long NumLiterals=0,NumLengths=0;
  
  CurRawArrayPtr = RawArray;
  RawArrayPtrDone = RawArray + RawFileLen;
  
  for(i=0;i<3;i++) { *CurRawArrayPtr++ = Literals[NumLiterals++]; }
  
  while ( CurRawArrayPtr < RawArrayPtrDone ) {
  	Hash = HASHFUNC(CurRawArrayPtr[-1],CurRawArrayPtr[-2],CurRawArrayPtr[-3]);
  
  	p = Table[Hash]; Table[Hash] = CurRawArrayPtr;
  	if ( p ) {
  	  MatchLen = Lengths[NumLengths++];
  	  if ( MatchLen > 0 ) {
  		if ( MatchLen == MatchLenEscape ) {
  		  long ReadLen;
  		  while( (ReadLen = Lengths[NumLengths++]) == MatchLenEscape )
  		    MatchLen += MatchLenEscape;
  			MatchLen += ReadLen;
  		}
  		while(MatchLen--) *CurRawArrayPtr++ = *p++;
  	  }
  	}
    #ifndef LiteralAfterMatch
  	else MatchLen = 0;
	
  	if ( MatchLen == 0 )
  #endif
  	*CurRawArrayPtr++ = Literals[NumLiterals++];
  }
}