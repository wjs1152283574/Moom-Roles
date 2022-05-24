* Moom-Roles
---

* 实现功能
    * http:
        1. 超级用户创建
        2. 图片验证码（支持配置是否需要使用）
        3. 普通管理员（CURD）
        4. 角色（CURD）
        5. 权限（CURD）
        6. 用户角色分配及取消
        7. 用户权限分配及取消
        8. 登陆/刷新登陆
        9. 获取用户信息

    * rpc
        1. 验证用户权限


* 记录:
    1. 权限新增路由字段，可以根据路由直接在项目中间件完成权限验证









* 问题记录：
    * `ratelimit.Server()` 全局限流器未开发完成 2022-05-21：`panic: cgroup cpu init failed!err:=not implemented yet`
