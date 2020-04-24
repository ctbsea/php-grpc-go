<?php
/**
 * Created by PhpStorm.
 * User: chentengbin
 * Date: 2020-04-24
 * Time: 18:40
 */

require __DIR__ . '/vendor/autoload.php';

$client = new  \Helloword\HelloWordClient('192.168.8.167:50051', [
    'credentials' => \Grpc\ChannelCredentials::createInsecure()
]);

$req = new \Helloword\MsgReq([]);

$res = $client->hi($req)->wait();

list($reply, $status) = $res;
$data = $reply->getMsg();
var_dump($data);



$req = new \Helloword\HelloRequest();
$req->setAge(10);
$req->setName("chen");

$res = $client->say($req)->wait();

list($reply, $status) = $res;
$data = $reply->getMessage();
var_dump($data);die;

