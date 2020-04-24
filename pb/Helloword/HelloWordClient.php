<?php
namespace Helloword;

class HelloWordClient extends \Grpc\BaseStub{

    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    public function hi(MsgReq $argument, $metadata=[], $options=[]){
        return $this->_simpleRequest('/helloword.Greeter/Msg',
            $argument,
            ['\Helloword\MsgRep', 'decode'],
            $metadata, $options);
    }

    public function say(HelloRequest $argument, $metadata=[], $options=[]){
        //method proto文件里面 package.server/rpc func
        return $this->_simpleRequest('/helloword.Greeter/SayHello',
            $argument,
            ['\Helloword\HelloReply', 'decode'], //响应解析类
            $metadata, $options);
    }
}


