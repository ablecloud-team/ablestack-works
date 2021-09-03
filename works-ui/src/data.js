
export const listdata = [
  "Racing car sprays burning fuel into crowd.",
  "Japanese princess to wed commoner.",
  "Australian walks 100km after outback crash.",
  "Man charged over missing wedding girl.",
  "Los Angeles battles huge wildfires.",
];
export const VMDiskListColumns = [
  {
    dataIndex: "Name",
    key: "Name",
    slots: { customRender: "nameRender" },
    title: "Name",
    sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "",
    key: "action",
    dataIndex: "action",
    slots: { customRender: "actionRender" },
  },
  {
    title: "State",
    dataIndex: "State",
    key: "State",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Size",
    dataIndex: "Size",
    key: "Size",
    sorter: (a, b) => (a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Connected Desktop",
    dataIndex: "Conn",
    key: "Conn",
    sorter: (a, b) => (a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
];
export const VMDiskListData = JSON.parse(
  '[{"Name":"Datadisk1","State":"Allocated","Size":"50GB","Conn":"VM1","action":""},' +
    '{"Name":"Datadisk2","State":"Allocated","Size":"100GB","Conn":"VM2","action":""},' +
    '{"Name":"Datadisk3","State":"Allocated","Size":"200GB","Conn":"VM3","action":""}]'
);

export const VMListColumns = [
  {
    dataIndex: "Name",
    key: "Name",
    slots: { customRender: "nameRender" },
    title: "Name",
    sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "",
    key: "action",
    dataIndex: "action",
    slots: { customRender: "actionRender" },
  },
  {
    title: "State",
    dataIndex: "State",
    key: "State",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Workspace",
    dataIndex: "Workspace",
    key: "Workspace",
    sorter: (a, b) =>
      a.Workspace < b.Workspace ? -1 : a.Workspace > b.Workspace ? 1 : 0,
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "User",
    dataIndex: "User",
    key: "User",
    sorter: (a, b) => (a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Connected",
    dataIndex: "Conn",
    key: "Conn",
    sorter: (a, b) => (a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
];
export const VMListData = JSON.parse(
  '[{"Name":"VM1","State":"Running","User":"user01","Conn":"TRUE","Workspace":"test1"},{"Name":"VM2","State":"Stopped","User":"user03","Conn":"FALSE","Workspace":"test1"},{"Name":"VM3","State":"Running","User":"user02","Conn":"TRUE","Workspace":"test1"}]'
);

export const NWListColumns = [
  {
    dataIndex: "Name",
    key: "Name",
    slots: { customRender: "nameRender" },
    title: "Name",
    sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "",
    key: "action",
    dataIndex: "action",
    slots: { customRender: "actionRender" },
  },
  {
    title: "State",
    dataIndex: "State",
    key: "State",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
];

export const UserListColumns = [
  {
    dataIndex: "Name",
    key: "Name",
    slots: { customRender: "nameRender" },
    title: "Name",
    sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  // {
  //     title: '',
  //     key: 'action',
  //     dataIndex: 'action',
  //     slots: {customRender: 'actionRender'}
  // },
  {
    title: "State",
    dataIndex: "State",
    key: "State",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Allocated Desktop",
    dataIndex: "Desktop",
    key: "Desktop",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
];

export const GroupPolicyColumns = [
  {
    dataIndex: "Name",
    key: "Name",
    slots: { customRender: "nameRender" },
    title: "Name",
    sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  // {
  //     title: '',
  //     key: 'action',
  //     dataIndex: 'action',
  //     slots: {customRender: 'actionRender'}
  // },
  {
    title: "State",
    dataIndex: "State",
    key: "State",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Allocated Desktop",
    dataIndex: "Desktop",
    key: "Desktop",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
];

export const AuditColumns = [
  {
    dataIndex: "Name",
    key: "Name",
    slots: { customRender: "nameRender" },
    title: "Name",
    sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  // {
  //     title: '',
  //     key: 'action',
  //     dataIndex: 'action',
  //     slots: {customRender: 'actionRender'}
  // },
  {
    title: "State",
    dataIndex: "State",
    key: "State",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
  {
    title: "Allocated Desktop",
    dataIndex: "Desktop",
    key: "Desktop",
    sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
    sortDirections: ["descend", "ascend"],
  },
];

export const NWListData = JSON.parse(
  '[{"Name":"Network01","State":"Allocated"}]'
);

export const UserListData = JSON.parse(
  '[{"Name":"user01","State":"Allocated","Desktop":"Desktop1"}]'
);

export const GroupPolicyData = JSON.parse(
  '[{"Name":"user01","State":"Allocated","Desktop":"Desktop1"}]'
);
export const AuditData = JSON.parse(
  '[{"Name":"user01","State":"Allocated","Desktop":"Desktop1"}]'
);
