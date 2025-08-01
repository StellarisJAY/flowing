import { defineStore } from 'pinia';
import { queryMenuTree } from '@/views/system/menu/api.js';
import { message } from 'ant-design-vue';

export const useMenuStore = defineStore('menuList', {
  state: () => {
    return {
      menuList: [],
      queryForm: {
        menuName: ""
      },
      menuForm: {
        menuName: '',
        path: '',
        component: '',
        parentId: 0,
        orderNum: 1,
        type: 1,
        icon: '',
        actionCode: '',
      },
      menuRules: {
        menuName: [
          {
            required: true,
            message: "请输入菜单名称",
            trigger: "submit"
          }
        ],
      },
      columns: [
        {
          title: '菜单名称',
          dataIndex: 'menuName',
          key: 'menuName',
        },
        {
          title: '类型',
          dataIndex: 'type',
          key: 'type',
        },
        {
          title: '路径',
          dataIndex: 'path',
          key: 'path',
        },
        {
          title: '组件',
          dataIndex: 'component',
          key: 'component',
        },
        {
          title: '操作',
          dataIndex: 'action',
          key: 'action',
        }
      ],
    };
  },
  actions: {
    async queryMenuList() {
      try {
        this.menuList = await queryMenuTree(this.queryForm);
      } catch {
        message.error('获取菜单列表失败');
      }
    },
    getColumns() {
      return this.columns;
    },
    getMenuTypeName(menuType) {
      switch (menuType) {
        case 1: return '目录';
        case 2: return '菜单';
        case 3: return '按钮';
        default: return '';
      }
    },
    getMenuTree() {
      return this.menuList;
    },
    clearQueryForm() {
      this.queryForm.menuName =  "";
    },
    getMenuTypeOptions() {
      return [
        {
          label: '目录',
          value: 1,
        },
        {
          label: '菜单',
          value: 2,
        },
        {
          label: '按钮',
          value: 3,
        },
      ];
    },
  },
});
