* Moom-Roles
---

#### 以下（HTTP）功能仅支持超级管理员

* 实现功能
    * http:
        * 用户功能
            * 图片验证码
            * 生成超级管理员
            * 用户登陆
            * 用户创建
            * 用户删除
            * 用户详细
            * 用户列表
            * 用户分配权限
            * 用户分配角色
            * 用户解除角色
            * 用户解除权限
            * 编辑基本信息
            * 用户锁定
            * 用户解锁
        * 角色功能
            * 角色创建
            * 角色删除
            * 角色列表
            * 角色编辑
        * 权限功能
            * 权限创建
            * 权限删除
            * 权限列表
            * 权限编辑
        * 路由功能
            * 路由创建
            * 路由删除
            * 路由列表
            * 路由编辑
            * 路由设置角色
            * 路由解除角色
            * 路由设置权限
            * 路由解除权限

    * rpc
        * 用户权限验证功能
            * 验证用户角色
            * 验证用户权限
        * 路由权限验证功能
            * 验证路由角色（Token）
            * 验证路由角色（UID）
            * 验证路由权限（Token）
            * 验证路由权限（UID）



* 记录:
    1. 权限新增路由字段，可以根据路由直接在项目中间件完成权限验证



#### TODO:
1. 后台功能需要限定超级用户可用；
2. 检测权限需要先查看用户是否依然存在并可用
3. 超管锁定后台用户





* 问题记录：
    * `ratelimit.Server()` 全局限流器未开发完成 2022-05-21：`panic: cgroup cpu init failed!err:=not implemented yet`
