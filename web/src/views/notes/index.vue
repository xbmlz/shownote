<template>
  <multipane class="custom-resizer" layout="vertical">
    <div class="pane menu">
      <div class="avatar">
        <a-avatar shape="square" size="large" :src="userInfo.avatar_url"></a-avatar>
        <h2>{{ userInfo.name }}</h2>
        <a-button type="primary" @click="createNote">
          新建笔记
          <PlusOutlined/>
        </a-button>
      </div>
      <div>
        <p>近期笔记</p>
        <p>我的分享</p>
        <a-directory-tree :tree-data="workspace.note"
                          v-model:selectedKeys="selectedKeys"
                          :replace-fields="replaceFields"
        >
          <template #title="node">
            <a-dropdown :trigger="['contextmenu']">

              <a-input v-if="node.isEdit" :value="node.name" @keyup.enter="updateName(node)"></a-input>
              <span v-else>{{ node.name }}</span>
              <template #overlay>
                <a-menu @click="({ key: menuKey }) => onContextMenuClick(node, menuKey)">
                  <a-menu-item key="1">新建子文件夹</a-menu-item>
                  <a-menu-item key="2">重命名</a-menu-item>
                  <a-menu-item key="3">删除</a-menu-item>
                  <a-menu-item key="4">向上移动</a-menu-item>
                  <a-menu-item key="5">向下移动</a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </template>
        </a-directory-tree>
      </div>
    </div>
    <multipane-resizer></multipane-resizer>
    <div class="pane catalogue">
      <div class="title">
        <a-input>
          <template #prefix>
            <search-outlined/>
          </template>
        </a-input>
        <a-divider></a-divider>
        <div class="info">
          <p>0篇笔记</p>
          <a-dropdown placement="bottomLeft">
            <UnorderedListOutlined/>
            <template #overlay>
              <a-menu>
                <a-menu-item>创建时间</a-menu-item>
                <a-menu-item>修改时间</a-menu-item>
                <a-menu-item>笔记名称</a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>
        <a-divider></a-divider>
        <a-empty :description="false">
          <a-button type="primary" @click="createNode">新建笔记</a-button>
        </a-empty>
      </div>
    </div>
    <multipane-resizer></multipane-resizer>
    <div class="pane doc" :style="{ flexGrow: 1 }">
      <div class="doc-title">
        <a-input></a-input>
        <a-button>保存</a-button>
      </div>
      <div id="vditor"></div>
    </div>
  </multipane>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, reactive, toRefs} from "vue";
import {useRouter, useRoute} from "vue-router"
import service from "@/utils/request.ts"
import {SearchOutlined, UnorderedListOutlined, PlusOutlined} from '@ant-design/icons-vue';
import {Multipane, MultipaneResizer} from "@/components/Multipane";
import {getNowDate} from "@/utils/index.ts";
import Vditor from "vditor";
import {message} from "ant-design-vue";


export default defineComponent({
  name: "App",
  components: {Multipane, MultipaneResizer, SearchOutlined, UnorderedListOutlined, PlusOutlined},
  setup() {
    const router = useRouter();
    const route = useRoute();
    const token = route.query.token || '';
    if (!token) {
      router.push('/login')
      return false
    }
    window.localStorage.setItem('token', token);

    let data = reactive({
      userInfo: {
        avatar_url: ""
      },
      workspace: {
        note: [],
        trash: [],
        share: []
      },
      workspaceSha: ''
    });

    onMounted(() => {
      new Vditor("vditor", {
        height: '100%',
        toolbarConfig: {
          pin: true,
        },
        cache: {
          enable: false,
        }
      });
    });
    //  初始化
    const init = async (): Promise<void> => {
      const info = await service.get(`/user/info?token=${token}`)
      data.userInfo = info.data;
      window.localStorage.setItem('userInfo', JSON.stringify(info.data));
      epoContent('.shownote/workspace.json').then(res => {
        let workspace = JSON.parse(res.content)
        console.log(workspace)
        data.workspaceSha = res.sha;
        data.workspace = workspace;
      })
    }

    const createNote = () => {

    }

    // 获取仓库具体路径下的内容
    function epoContent(path: any) {
      return new Promise<Pro>(async resolve => {
        const token = localStorage.token
        const info = JSON.parse(localStorage.userInfo)
        const res = await service.get('/repo/content', {
          params: {
            token: token,
            login: info.login,
            path: path
          }
        })
        resolve(res.data)
      })
    }

    //  更新目录JSON
    const updateWorkspace = async (content: Array<[]>, sha: string) => {
      const token = localStorage.token
      const userInfo = JSON.parse(localStorage.userInfo)
      const res = await service.put('/repo/file', {
        "content": JSON.stringify(content),
        "login": userInfo.login,
        "path": '.shownote/workspace.json',
        "sha": sha,
        "token": token
      })
      //data.workspace = JSON.parse(res.data.content);
      data.workspaceSha = res.data.sha;
    }
    // 更新文件
    const updateFile = async (content: any, path: string, sha: string) => {
      const token = localStorage.token
      const userInfo = JSON.parse(localStorage.userInfo)
      await service.put('/repo/file', {
        "content": JSON.stringify(content),
        "login": userInfo.login,
        "path": path,
        "sha": sha,
        "token": token
      })
    }
    init()

    const selectedKeys = ref<string[]>([]);
    const replaceFields = {
      children: 'child',
      title: 'name',
      key: 'path'
    }
    const onContextMenuClick = (node: object, menuKey: string) => {
      console.log(`treeKey: ${node.path}, menuKey: ${menuKey}`);
      //  新建子文件夹
      if (menuKey === '1') {
        node.child.unshift({
          "name": "新文件夹",
          "path": `${node.path}/新文件夹`,
          "isDir": true,
          "isShare": false,
          "sha": "",
          "size": 0,
          "url": "",
          "type": "",
          "html_url": "",
          "download_url": "",
          "createTime": getNowDate(),
          "updateTime": getNowDate(),
          "child": []
        })
        updateWorkspace(data.workspace, data.workspaceSha)
      }
      if (menuKey === '2') {
        // {
        //   "name": "新文件夹",
        //     "path": `${treeKey}/新文件夹`,
        //     "updateTime": getNowDate(),
        // }
      }
      //  删除
          if (menuKey === '3') {
            if(node.child.length) {
              return message.error('存在文件，不允许删除')
            }
            console.log(node)
          }

    };
    const updateName = node => {

    }

    return {
      ...toRefs(data),
      selectedKeys,
      replaceFields,
      onContextMenuClick,
      createNote,
      updateName
    }
  },
});
</script>

<style scoped lang="scss">
@import url(vditor/dist/index.css);

.custom-resizer {
  width: 100%;
  height: 100vh;
}

.custom-resizer > .pane {
  text-align: left;
  padding: 15px;
  overflow: hidden;
  background: #fff;
  border: 1px solid #ccc;

  &.menu {
    width: 20%;
    max-width: 250px;
  }

  &.catalogue {
    width: 300px;
    background-color: #f4f6f9;
  }
}

.custom-resizer > .multipane-resizer {
  margin: 0;
  left: 0;
  position: relative;

  &:before {
    display: block;
    content: "";
    width: 3px;
    height: 40px;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -20px;
    margin-left: -1.5px;
    border-left: 1px solid #ccc;
    border-right: 1px solid #ccc;
  }

  &:hover {
    &:before {
      border-color: #999;
    }
  }
}

.menu {
  .avatar {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin-bottom: 20px;
  }
}

.catalogue {
  .ant-divider-horizontal {
    margin: 15px 0;
  }

  .info {
    display: flex;
    align-items: center;
    justify-content: space-between;

    p {
      margin: 0;
    }
  }
}

.doc {
  &-title {
    display: flex;
    margin-bottom: 10px;
  }
}
</style>