
export const columns = [
    {
        dataIndex: 'Name',
        key: 'Name',
        slots: {customRender: 'nameRender'},
        title: 'Name',
        sorter: (a, b) => a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0,
        sortDirections: ['descend', 'ascend']
    },
    {
        title:'',
        key:'action',
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
export const listdata = [
    'Racing car sprays burning fuel into crowd.',
    'Japanese princess to wed commoner.',
    'Australian walks 100km after outback crash.',
    'Man charged over missing wedding girl.',
    'Los Angeles battles huge wildfires.',
];
export const data2=JSON.parse(
    '[{"Name":"ubuntu-desktop","Type":"Desktop","State":"Running","IP Address":"10.1.1.17","Account":"djpark","Zone":"zone","tag":["nice","developer"],"action":"","NoC":42,"NoD":83,"NetType":"Isolated"},{"Name":"guacamole-extension","Type":"App","State":"Stopped","IP Address":"10.1.1.192","Account":"djpark","Zone":"zone","tag":["loser"],"action":"","NoC":19,"NoD":"","NetType":"Shared"},{"Name":"sicho-ansible-host2","Type":"Desktop","State":"Running","IP Address":"10.1.1.160","Account":"sicho","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":55,"NoD":81,"NetType":"L2"},{"Name":"sicho-ansible-host1","Type":"App","State":"Running","IP Address":"10.1.1.230","Account":"sicho","Zone":"zone","tag":["nice","developer"],"action":"","NoC":69,"NoD":"","NetType":"L2"},{"Name":"sicho-ansible-ctrlSvr","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.54","Account":"sicho","Zone":"zone","tag":["loser"],"action":"","NoC":99,"NoD":18,"NetType":"Shared"},{"Name":"ycyun-dc-test","Type":"App","State":"Running","IP Address":"10.1.1.12","Account":"ycyun","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":34,"NoD":"","NetType":"Isolated"},{"Name":"ycyun-server-2019","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.70","Account":"ycyun","Zone":"zone","tag":["nice","developer"],"action":"","NoC":23,"NoD":90,"NetType":"Shared"},{"Name":"hwryu-centos-dev-01","Type":"App","State":"Running","IP Address":"10.1.1.35","Account":"hwryu","Zone":"zone","tag":["loser"],"action":"","NoC":67,"NoD":"","NetType":"L2"},{"Name":"guacamole-docker","Type":"Desktop","State":"Running","IP Address":"10.1.1.9","Account":"djpark","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":54,"NoD":2,"NetType":"L2"},{"Name":"wallvm-tj","Type":"App","State":"Stopped","IP Address":"10.1.1.166","Account":"tjbae","Zone":"zone","tag":["nice","developer"],"action":"","NoC":55,"NoD":"","NetType":"Shared"},{"Name":"smlee-dev","Type":"Desktop","State":"Running","IP Address":"192.168.0.147","Account":"smlee","Zone":"zone","tag":["loser"],"action":"","NoC":84,"NoD":80,"NetType":"Isolated"},{"Name":"docs-server-01","Type":"App","State":"Stopped","IP Address":"10.1.1.178","Account":"sicho","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":8,"NoD":"","NetType":"Shared"},{"Name":"dhs-docs-dev-svr","Type":"Desktop","State":"Running","IP Address":"","Account":"admin","Zone":"zone","tag":["nice","developer"],"action":"","NoC":7,"NoD":2,"NetType":"L2"},{"Name":"dhs-win10","Type":"App","State":"Running","IP Address":"","Account":"admin","Zone":"zone","tag":["loser"],"action":"","NoC":87,"NoD":"","NetType":"L2"},{"Name":"mold-install-test","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.154","Account":"djpark","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":3,"NoD":36,"NetType":"Shared"},{"Name":"ktcheon-centos8-min","Type":"App","State":"Running","IP Address":"10.1.1.225","Account":"ktcheon","Zone":"zone","tag":["nice","developer"],"action":"","NoC":52,"NoD":"","NetType":"Isolated"},{"Name":"ktcheon-windows10","Type":"Desktop","State":"Stopped","IP Address":"10.1.1.107","Account":"ktcheon","Zone":"zone","tag":["loser"],"action":"","NoC":73,"NoD":56,"NetType":"Shared"},{"Name":"home-dev-svr","Type":"App","State":"Running","IP Address":"","Account":"admin","Zone":"zone","tag":["cool","teacher"],"action":"","NoC":62,"NoD":"","NetType":"L2"},{"Name":"ktcheon-centos8","Type":"Desktop","State":"Running","IP Address":"10.1.1.17","Account":"ktcheon","Zone":"zone","tag":["nice","developer"],"action":"","NoC":66,"NoD":77,"NetType":"L2"},{"Name":"mold-build-server","Type":"App","State":"Stopped","IP Address":"10.1.1.107","Account":"djpark","Zone":"zone","tag":["loser"],"action":"","NoC":70,"NoD":"","NetType":"Shared"}]'
)
export const data = [
    {
        key: 1,
        Display_Name: "ubuntu-desktop",
        Type: "Desktop",
        Display_Name1: "ubuntu-desktop",
        State: "Running",
        IPAddress: "10.1.1.17",
        Account: "djpark",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 2,
        Display_Name: "guacamole-extension",
        Type: "Apps",
        Display_Name1: "guacamole-extension",
        State: "Stopped",
        IPAddress: "10.1.1.192",
        Account: "djpark",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
    {
        key: 3,
        Display_Name: "sicho-ansible-host2",
        Type: "Desktop",
        Display_Name1: "sicho-ansible-host2",
        State: "Running",
        IPAddress: "10.1.1.160",
        Account: "sicho",
        Zone: "zone",
        tag: [
            "cool",
            "teacher"
        ]
    },
    {
        key: 4,
        Display_Name: "sicho-ansible-host1",
        Type: "Desktop",
        Display_Name1: "sicho-ansible-host1",
        State: "Running",
        IPAddress: "10.1.1.230",
        Account: "sicho",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 5,
        Display_Name: "sicho-ansible-ctrlSvr",
        Type: "Desktop",
        Display_Name1: "sicho-ansible-ctrlSvr",
        State: "Running",
        IPAddress: "10.1.1.54",
        Account: "sicho",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
    {
        key: 6,
        Display_Name: "ycyun-dc-test",
        Type: "Apps",
        Display_Name1: "ycyun-dc-test",
        State: "Running",
        IPAddress: "10.1.1.12",
        Account: "ycyun",
        Zone: "zone",
        tag: [
            "cool",
            "teacher"
        ]
    },
    {
        key: 7,
        Display_Name: "ycyun-server-2019",
        Type: "Desktop",
        Display_Name1: "ycyun-server-2019",
        State: "Running",
        IPAddress: "10.1.1.70",
        Account: "ycyun",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 8,
        Display_Name: "hwryu-centos-dev-01",
        Type: "Desktop",
        Display_Name1: "hwryu-centos-dev-01",
        State: "Running",
        IPAddress: "10.1.1.35",
        Account: "hwryu",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
    {
        key: 9,
        Display_Name: "guacamole-docker",
        Type: "Desktop",
        Display_Name1: "guacamole-docker",
        State: "Running",
        IPAddress: "10.1.1.9",
        Account: "djpark",
        Zone: "zone",
        tag: [
            "cool",
            "teacher"
        ]
    },
    {
        key: 10,
        Display_Name: "wallvm-tj",
        Type: "Desktop",
        Display_Name1: "wallvm-tj",
        State: "Running",
        IPAddress: "10.1.1.166",
        Account: "tjbae",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 11,
        Display_Name: "smlee-dev",
        Type: "Desktop",
        Display_Name1: "smlee-dev",
        State: "Running",
        IPAddress: "192.168.0.147",
        Account: "smlee",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
    {
        key: 12,
        Display_Name: "docs-server-01",
        Type: "Desktop",
        Display_Name1: "docs-server-01",
        State: "Running",
        IPAddress: "10.1.1.178",
        Account: "sicho",
        Zone: "zone",
        tag: [
            "cool",
            "teacher"
        ]
    },
    {
        key: 13,
        Display_Name: "dhs-docs-dev-svr",
        Type: "Desktop",
        Display_Name1: "dhs-docs-dev-svr",
        State: "Running",
        IPAddress: "",
        Account: "admin",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 14,
        Display_Name: "dhs-win10",
        Type: "Desktop",
        Display_Name1: "dhs-win10",
        State: "Running",
        IPAddress: "",
        Account: "admin",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
    {
        key: 15,
        Display_Name: "mold-install-test",
        Type: "Desktop",
        Display_Name1: "mold-install-test",
        State: "Running",
        IPAddress: "10.1.1.154",
        Account: "djpark",
        Zone: "zone",
        tag: [
            "cool",
            "teacher"
        ]
    },
    {
        key: 16,
        Display_Name: "ktcheon-centos8-min",
        Type: "Desktop",
        Display_Name1: "ktcheon-centos8-min",
        State: "Running",
        IPAddress: "10.1.1.225",
        Account: "ktcheon",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 17,
        Display_Name: "ktcheon-windows10",
        Type: "Desktop",
        Display_Name1: "ktcheon-windows10",
        State: "Running",
        IPAddress: "10.1.1.107",
        Account: "ktcheon",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
    {
        key: 18,
        Display_Name: "home-dev-svr",
        Type: "Desktop",
        Display_Name1: "home-dev-svr",
        State: "Running",
        IPAddress: "",
        Account: "admin",
        Zone: "zone",
        tag: [
            "cool",
            "teacher"
        ]
    },
    {
        key: 19,
        Display_Name: "ktcheon-centos8",
        Type: "Desktop",
        Display_Name1: "ktcheon-centos8",
        State: "Running",
        IPAddress: "10.1.1.17",
        Account: "ktcheon",
        Zone: "zone",
        tag: [
            "nice",
            "developer"
        ]
    },
    {
        key: 20,
        Name: "mold-build-server",
        Type: "Desktop",
        Display_Name1: "mold-build-server",
        State: "Running",
        IPAddress: "10.1.1.107",
        Account: "djpark",
        Zone: "zone",
        tag: [
            "loser"
        ]
    },
];