export const WorkspaceListColumns = [
    {
        dataIndex: 'Name',
        key: 'Name',
        slots: {customRender: 'nameRender'},
        title: 'Name',
        sorter: (a, b) => a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: '',
        key: 'action',
        dataIndex: 'action',
        slots: {customRender: 'actionRender'}
    },
    {
        title: 'State',
        dataIndex: 'State',
        key: 'State',
        sorter: (a, b) => a.State < b.State ? -1 : a.State > b.State ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Type',
        dataIndex: 'Type',
        key: 'Type',
        sorter: (a, b) => a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Number Of Desktop',
        dataIndex: 'NoD',
        key: 'Nod',
        sorter: (a, b) => a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Number Of Connection',
        dataIndex: 'NoC',
        key: 'NoC',
        sorter: (a, b) => a.NoC < b.NoC ? -1 : a.NoC > b.NoC ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Network Type',
        dataIndex: 'NetType',
        key: 'NetType',
        sorter: (a, b) => a.NetType < b.NetType ? -1 : a.NetType > b.NetType ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Restrict',
        dataIndex: 'Restrict',
        key: 'Restrict',
        sorter: (a, b) => a.Restrict < b.Restrict ? -1 : a.Restrict > b.Restrict ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    }
    // ,
    // {
    //     title: 'Tags',
    //     key: 'tags',
    //     dataIndex: 'tag',
    //     slots: {customRender: 'tags'},
    // }
];
export const WorkspaceListData = JSON.parse(
    '[{"Name":"ubuntu-desktop","Type":"Desktop","State":"Running","IP Address":"10.1.1.17","Account":"djpark","Zone":"zone","tag":["nice","developer"],"action":"","NoC":42,"NoD":83,"NetType":"Isolated"},{"Name":"guacamole-extension","Type":"App","State":"Stopped","IP Address":"10.1.1.192","Account":"djpark","Zone":"zone","tag":["loser"],"action":"","NoC":19,"NoD":"","NetType":"Shared"},{"Name":"sicho-ansible-host2","Type":"Desktop","State":"Running","IP Address":"10.1.1.160","Account":"sicho","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":55,"NoD":81,"NetType":"L2"},{"Name":"sicho-ansible-host1","Type":"App","State":"Running","IP Address":"10.1.1.230","Account":"sicho","Zone":"zone","tag":["nice","developer"],"action":"","NoC":69,"NoD":"","NetType":"L2"},{"Name":"sicho-ansible-ctrlSvr","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.54","Account":"sicho","Zone":"zone","tag":["loser"],"action":"","NoC":99,"NoD":18,"NetType":"Shared"},{"Name":"ycyun-dc-test","Type":"App","State":"Running","IP Address":"10.1.1.12","Account":"ycyun","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":34,"NoD":"","NetType":"Isolated"},{"Name":"ycyun-server-2019","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.70","Account":"ycyun","Zone":"zone","tag":["nice","developer"],"action":"","NoC":23,"NoD":90,"NetType":"Shared"},{"Name":"hwryu-centos-dev-01","Type":"App","State":"Running","IP Address":"10.1.1.35","Account":"hwryu","Zone":"zone","tag":["loser"],"action":"","NoC":67,"NoD":"","NetType":"L2"},{"Name":"guacamole-docker","Type":"Desktop","State":"Running","IP Address":"10.1.1.9","Account":"djpark","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":54,"NoD":2,"NetType":"L2"},{"Name":"wallvm-tj","Type":"App","State":"Stopped","IP Address":"10.1.1.166","Account":"tjbae","Zone":"zone","tag":["nice","developer"],"action":"","NoC":55,"NoD":"","NetType":"Shared"},{"Name":"smlee-dev","Type":"Desktop","State":"Running","IP Address":"192.168.0.147","Account":"smlee","Zone":"zone","tag":["loser"],"action":"","NoC":84,"NoD":80,"NetType":"Isolated"},{"Name":"docs-server-01","Type":"App","State":"Stopped","IP Address":"10.1.1.178","Account":"sicho","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":8,"NoD":"","NetType":"Shared"},{"Name":"dhs-docs-dev-svr","Type":"Desktop","State":"Running","IP Address":"","Account":"admin","Zone":"zone","tag":["nice","developer"],"action":"","NoC":7,"NoD":2,"NetType":"L2"},{"Name":"dhs-win10","Type":"App","State":"Running","IP Address":"","Account":"admin","Zone":"zone","tag":["loser"],"action":"","NoC":87,"NoD":"","NetType":"L2"},{"Name":"mold-install-test","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.154","Account":"djpark","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":3,"NoD":36,"NetType":"Shared"},{"Name":"ktcheon-centos8-min","Type":"App","State":"Running","IP Address":"10.1.1.225","Account":"ktcheon","Zone":"zone","tag":["nice","developer"],"action":"","NoC":52,"NoD":"","NetType":"Isolated"},{"Name":"ktcheon-windows10","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.107","Account":"ktcheon","Zone":"zone","tag":["loser"],"action":"","NoC":73,"NoD":56,"NetType":"Shared"},{"Name":"home-dev-svr","Type":"App","State":"Running","IP Address":"","Account":"admin","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":62,"NoD":"","NetType":"L2"},{"Name":"ktcheon-centos8","Type":"Desktop","State":"Running","IP Address":"10.1.1.17","Account":"ktcheon","Zone":"zone","tag":["nice","developer"],"action":"","NoC":66,"NoD":77,"NetType":"L2"},{"Name":"mold-build-server","Type":"App","State":"Stopped","IP Address":"10.1.1.107","Account":"djpark","Zone":"zone","tag":["loser"],"action":"","NoC":70,"NoD":"","NetType":"Shared"}]'
)
export const listdata = [
    'Racing car sprays burning fuel into crowd.',
    'Japanese princess to wed commoner.',
    'Australian walks 100km after outback crash.',
    'Man charged over missing wedding girl.',
    'Los Angeles battles huge wildfires.',
];
export const VMDiskListColumns = [
    {
        dataIndex: 'Name',
        key: 'Name',
        slots: {customRender: 'nameRender'},
        title: 'Name',
        sorter: (a, b) => a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: '',
        key: 'action',
        dataIndex: 'action',
        slots: {customRender: 'actionRender'}
    },
    {
        title: 'State',
        dataIndex: 'State',
        key: 'State',
        sorter: (a, b) => a.State < b.State ? -1 : a.State > b.State ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Size',
        dataIndex: 'Size',
        key: 'Size',
        sorter: (a, b) => a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Connected Desktop',
        dataIndex: 'Conn',
        key: 'Conn',
        sorter: (a, b) => a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    }
];
export const VMDiskListData = JSON.parse(
    '[{"Name":"Datadisk1","State":"Allocated","Size":"50GB","Conn":"TRUE","action":""},{"Name":"Datadisk2","State":"Allocated","Size":"100GB","Conn":"FALSE","action":""},{"Name":"Datadisk3","State":"Allocated","Size":"200GB","Conn":"TRUE","action":""}]'
)

export const VMListColumns = [
    {
        dataIndex: 'Name',
        key: 'Name',
        slots: {customRender: 'nameRender'},
        title: 'Name',
        sorter: (a, b) => a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: '',
        key: 'action',
        dataIndex: 'action',
        slots: {customRender: 'actionRender'}
    },
    {
        title: 'State',
        dataIndex: 'State',
        key: 'State',
        sorter: (a, b) => a.State < b.State ? -1 : a.State > b.State ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Workspace',
        dataIndex: 'Workspace',
        key: 'Workspace',
        sorter: (a, b) => a.Workspace < b.Workspace ? -1 : a.Workspace > b.Workspace ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'User',
        dataIndex: 'User',
        key: 'User',
        sorter: (a, b) => a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: 'Connected',
        dataIndex: 'Conn',
        key: 'Conn',
        sorter: (a, b) => a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    }
];
export const VMListData = JSON.parse(
    '[{"Name":"VM1","State":"Running","User":"user01","Conn":"TRUE","Workspace":"test1"},{"Name":"VM2","State":"Stopped","User":"user03","Conn":"FALSE","Workspace":"test1"},{"Name":"VM3","State":"Running","User":"user02","Conn":"TRUE","Workspace":"test1"}]'
)

export const NWListColumns = [
    {
        dataIndex: 'Name',
        key: 'Name',
        slots: {customRender: 'nameRender'},
        title: 'Name',
        sorter: (a, b) => a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title: '',
        key: 'action',
        dataIndex: 'action',
        slots: {customRender: 'actionRender'}
    },
    {
        title: 'State',
        dataIndex: 'State',
        key: 'State',
        sorter: (a, b) => a.State < b.State ? -1 : a.State > b.State ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    ];
export const NWListData = JSON.parse(
    '[{"Name":"Network01","State":"Allocated"}]'
)