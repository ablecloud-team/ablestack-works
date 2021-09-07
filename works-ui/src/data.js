
export const listdata = [
  "Racing car sprays burning fuel into crowd.",
  "Japanese princess to wed commoner.",
  "Australian walks 100km after outback crash.",
  "Man charged over missing wedding girl.",
  "Los Angeles battles huge wildfires.",
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

export const GroupPolicyData = JSON.parse(
  '[{"Name":"GroupPolicy-01","State":"Allocated","Desktop":"Desktop1"}]'
);

export const UserListData = JSON.parse(
  '[{"Name":"user01","State":"Allocated","Desktop":"Desktop1"}]'
);
export const AuditData = JSON.parse(
  '[{"Name":"user01","State":"Allocated","Desktop":"Desktop1"}]'
);
