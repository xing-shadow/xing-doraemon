const Router = [
    {
        title: '数据源管理',
        key: '/antd/dist/prom',
        subs:[
            {
                title: '数据源列表',
                key: '/antd/dist/prom/list',
            },
            {
                title: '数据源添加',
                key: '/antd/dist/prom/add',
            }
        ]
    },
    {
        title: '报警计划管理',
        key: '/antd/dist/plan',
        subs:[
            {
                title: '报警计划列表',
                key: '/antd/dist/plan/list',
            }
        ]
    },
    {
        title: '告警规则管理',
        key: '/antd/dist/rule',
        subs:[
            {
                title: '告警规则列表',
                key: '/antd/dist/rule/list',
            }
        ]
    },
    {
        title: '告警信息管理',
        key: '/antd/dist/alert/list',
        subs:[
            {
                title: '告警信息列表',
                key: '/antd/dist/alert/list',
            }
        ]
    },
]

export default Router;