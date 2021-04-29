<?php
/** 
 * @OA\Info(title="My First API", version="0.1")
 */
class Customer
{
    /**
     * @OA\Get(
     *     path="/customer/info",
     *     summary="用户的个人信息",
     *     description="这不是个api接口,这个返回一个页面",
     *     @OA\Parameter(name="userId", in="query", @OA\Schema(type="string"), required=true, description="用户ID"),
     *     @OA\Response(
     *      response="200",
     *      description="An example resource"
     *     )
     * )
     */
    public function info(int $userId, string $userToken)
    {

    }
}