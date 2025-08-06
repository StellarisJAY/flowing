import { defineStore } from 'pinia';
import { createMenu, queryMenuTree } from '@/views/system/menu/api.js';
import { message } from 'ant-design-vue';

export const useMenuStore = defineStore('menuList', {
  state: () => {
    return {
      menuList: [],
      parentMenuOptions: [],
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
        },
      ],
      menuTypeOptions: [
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
      ],
      menuFormRules: {
        menuName: [{ required: true, message: '请输入菜单名称', trigger: 'submit' }],
        path: [{ required: true, message: '请输入路径', trigger: 'submit' }],
        parentId: [{ required: true, message: '请选择父菜单', trigger: 'submit' }],
        actionCode: [{ required: true, message: '请输入操作权限', trigger: 'submit' }],
      },
      queryFormSchema: [
        {
          name: 'menuName',
          label: '菜单名称',
          type: 'input',
          placeholder: '请输入菜单名称',
          defaultValue: '',
        }
      ]
    };
  },
  actions: {
    async addMenu(menu) {
      try {
        await createMenu(menu);
        message.success('添加成功');
        return true;
      }catch (err) {
        message.error(err);
        return false;
      }
    },
    async queryMenuList(query) {
      try {
        const res = await queryMenuTree(query);
        this.menuList = res.data;
      } catch {
        message.error('获取菜单列表失败');
      }
    },
    getMenuTypeName(menuType) {
      switch (menuType) {
        case 1:
          return '目录';
        case 2:
          return '菜单';
        case 3:
          return '按钮';
        default:
          return '';
      }
    },
    async getParentMenuOptions() {
      try {
        const res = await queryMenuTree({});
        const menuList = res.data;
        const buildOptions = (menus) => {
          const options = menus
            .filter((item) => item.type !== 3)
            .map((item) => {
              return {
                label: item.menuName,
                value: item.id,
                children: item.children,
              };
            });
          for (const item of options) {
            if (item.children) {
              item.children = buildOptions(item.children);
            }
          }
          return options;
        };

        const options = buildOptions(menuList);
        options.unshift({
          label: '根目录',
          value: 0,
        });
        this.parentMenuOptions = options;
      } catch {
        message.error('获取父菜单列表失败');
      }
    },
    getMenuFormSchema() {
      return [
        {
          name: 'menuName',
          label: '菜单名称',
          type: 'input',
          placeholder: '请输入菜单名称',
          defaultValue: '',
        },
        {
          name: 'type',
          label: '类型',
          type: 'radioGroup',
          placeholder: '请选择类型',
          defaultValue: 1,
          options: () => this.menuTypeOptions,
        },
        {
          name: 'path',
          label: '路径',
          type: 'input',
          placeholder: '请输入路径',
          defaultValue: '',
        },
        {
          name: 'component',
          label: '组件',
          type: 'input',
          placeholder: '请输入前端组件路径',
          defaultValue: '',
        },
        {
          name: 'parentId',
          label: '父菜单',
          type: 'treeSelect',
          placeholder: '请选择父菜单',
          defaultValue: 0,
          options: () => this.parentMenuOptions,
        },
        {
          name: 'orderNum',
          label: '排序',
          type: 'inputNumber',
          placeholder: '请输入排序',
          defaultValue: 1,
        },
        {
          name: 'icon',
          label: '图标',
          type: 'input',
          placeholder: '请输入图标',
          defaultValue: '',
        },
        {
          name: 'actionCode',
          label: '权限标识',
          type: 'input',
          placeholder: '请输入权限标识',
          defaultValue: '',
          hidden: (formState) => formState.type !== 3,
        },
      ];
    },
    initMenuForm() {
      this.menuForm = {
        menuName: '',
        path: '',
        component: '',
        parentId: 0,
        orderNum: 1,
        type: 1,
        icon: '',
        actionCode: '',
      };
    },
  },
});
