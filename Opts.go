package anthropic

type Opts struct{
	Context []MessageModule
	ContextID string
	Sender Sender
}

func (o *Opts) Null(){
	
}