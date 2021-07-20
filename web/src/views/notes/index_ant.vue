<!--<template>-->
<!--  <multipane class="custom-resizer" layout="vertical">-->
<!--    <div class="pane menu">-->
<!--      <div class="avatar">-->
<!--        <a-avatar-->
<!--          shape="square"-->
<!--          size="large"-->
<!--          :src="userInfo.avatar_url"-->
<!--        ></a-avatar>-->
<!--        <h2>{{ userInfo.name }}</h2>-->
<!--        <a-button type="primary" @click="createNote">-->
<!--          新建笔记-->
<!--          <PlusOutlined />-->
<!--        </a-button>-->
<!--      </div>-->
<!--      <div>-->
<!--        <p>近期笔记</p>-->
<!--        <p>我的分享</p>-->
<!--        <a-directory-tree-->
<!--          :tree-data="workspaceNotes"-->
<!--          v-model:selectedKeys="selectedKeys"-->
<!--          :replace-fields="replaceFields"-->
<!--          @select="selectTreeNode"-->
<!--        >-->
<!--          <template #title="node">-->
<!--            <a-dropdown :trigger="['contextmenu']">-->
<!--              <a-input-->
<!--                v-if="node.isEdit"-->
<!--                :value="node.name"-->
<!--                @keyup.enter="updateName(node)"-->
<!--              ></a-input>-->
<!--              <span v-else>{{ node.name }}</span>-->
<!--              <template #overlay>-->
<!--                <a-menu-->
<!--                  @click="-->
<!--                    ({ key: menuKey }) => onContextMenuClick(node, menuKey)-->
<!--                  "-->
<!--                >-->
<!--                  <a-menu-item key="0">新建文件</a-menu-item>-->
<!--                  <a-menu-item key="1">新建文件夹</a-menu-item>-->
<!--                  &lt;!&ndash;                  <a-menu-item key="2">重命名</a-menu-item>&ndash;&gt;-->
<!--                  <a-menu-item key="3">测试清空目录</a-menu-item>-->
<!--                  &lt;!&ndash;                  <a-menu-item key="4">向上移动</a-menu-item>&ndash;&gt;-->
<!--                  &lt;!&ndash;                  <a-menu-item key="5">向下移动</a-menu-item>&ndash;&gt;-->
<!--                </a-menu>-->
<!--              </template>-->
<!--            </a-dropdown>-->
<!--          </template>-->
<!--        </a-directory-tree>-->
<!--      </div>-->
<!--    </div>-->
<!--    <multipane-resizer></multipane-resizer>-->
<!--    <div class="pane doc" :style="{ flexGrow: 1 }">-->
<!--      <tempalte v-show="activeTreeNode">-->
<!--        <div class="doc-title">-->
<!--          <a-input v-model:value="activeTreeNode.name"></a-input>-->
<!--          <a-button @click="updateNote">保存</a-button>-->
<!--        </div>-->
<!--        <div id="vditor"></div>-->
<!--      </tempalte>-->
<!--      <a-empty v-show="!activeTreeNode"></a-empty>-->
<!--    </div>-->
<!--  </multipane>-->
<!--</template>-->

<!--<script lang="ts">-->
<!--import { ref, defineComponent, onMounted, reactive, toRefs } from "vue";-->
<!--import { useRouter, useRoute } from "vue-router";-->
<!--import service from "@/utils/request";-->
<!--import {-->
<!--  SearchOutlined,-->
<!--  UnorderedListOutlined,-->
<!--  PlusOutlined,-->
<!--} from "@ant-design/icons-vue";-->
<!--import { Multipane, MultipaneResizer } from "@/components/Multipane";-->
<!--import { getNowDate } from "@/utils/date";-->
<!--import Vditor from "vditor";-->
<!--import { message } from "ant-design-vue";-->

<!--export default defineComponent({-->
<!--  name: "App",-->
<!--  components: {-->
<!--    Multipane,-->
<!--    MultipaneResizer,-->
<!--    SearchOutlined,-->
<!--    UnorderedListOutlined,-->
<!--    PlusOutlined,-->
<!--  },-->
<!--  setup() {-->
<!--    const router = useRouter();-->
<!--    const route = useRoute();-->
<!--    const token: any = route.query.token || "";-->
<!--    if (!token) {-->
<!--      router.push("/login");-->
<!--      return false;-->
<!--    }-->
<!--    window.localStorage.setItem("token", token);-->

<!--    let data: any = reactive({-->
<!--      vditor: null,-->
<!--      userInfo: {-->
<!--        avatar_url: "",-->
<!--      },-->
<!--      workspace: {-->
<!--        note: [],-->
<!--        trash: [],-->
<!--        share: [],-->
<!--      },-->
<!--      workspaceNotes: [-->
<!--        {-->
<!--          name: "我的笔记",-->
<!--          path: "/note",-->
<!--          isDir: true,-->
<!--          isShare: false,-->
<!--          sha: "",-->
<!--          size: 0,-->
<!--          url: "",-->
<!--          type: "",-->
<!--          html_url: "",-->
<!--          download_url: "",-->
<!--          createTime: "",-->
<!--          updateTime: "",-->
<!--          child: [],-->
<!--        },-->
<!--      ],-->
<!--      activeTreeNode: {-->
<!--        name: "",-->
<!--      }, // 当前选中的目录-->
<!--      workspaceSha: "",-->
<!--    });-->

<!--    onMounted(() => {-->
<!--      data.vditor = new Vditor("vditor", {-->
<!--        height: "100%",-->
<!--        toolbarConfig: {-->
<!--          pin: true,-->
<!--        },-->
<!--        cache: {-->
<!--          enable: false,-->
<!--        },-->
<!--      });-->
<!--    });-->
<!--    //  初始化-->
<!--    const init = async (): Promise<void> => {-->
<!--      const info = await service.get(`/user/info?token=${token}`);-->
<!--      data.userInfo = info.data;-->
<!--      window.localStorage.setItem("userInfo", JSON.stringify(info.data));-->
<!--      epoContent(".shownote/workspace.json").then((res: any) => {-->
<!--        let workspace = JSON.parse(res.content);-->
<!--        console.log(workspace, "workspace&#45;&#45;&#45;&#45;&#45;&#45;");-->
<!--        data.workspaceSha = res.sha;-->
<!--        data.workspace = workspace;-->
<!--        data.workspaceNotes[0].child = workspace.note;-->
<!--      });-->
<!--    };-->

<!--    // 获取仓库具体路径下的内容-->
<!--    function epoContent(path: any) {-->
<!--      return new Promise(async (resolve) => {-->
<!--        const token = localStorage.token;-->
<!--        const info = JSON.parse(localStorage.userInfo);-->
<!--        const res = await service.get("/repo/content", {-->
<!--          params: {-->
<!--            token: token,-->
<!--            login: info.login,-->
<!--            path: path,-->
<!--          },-->
<!--        });-->
<!--        resolve(res.data);-->
<!--      });-->
<!--    }-->

<!--    //  更新目录JSON-->
<!--    const updateWorkspace = async (content: Array<[]>, sha: string) => {-->
<!--      const token = localStorage.token;-->
<!--      const userInfo = JSON.parse(localStorage.userInfo);-->
<!--      const res = await service.put("/repo/file", {-->
<!--        content: JSON.stringify(content),-->
<!--        login: userInfo.login,-->
<!--        path: ".shownote/workspace.json",-->
<!--        sha: sha,-->
<!--        token: token,-->
<!--      });-->
<!--      //data.workspace = JSON.parse(res.data.content);-->
<!--      // data.workspaceNotes[0].child = JSON.parse(res.data.content);-->
<!--      data.workspaceSha = res.data.sha;-->
<!--    };-->
<!--    // 更新文件-->
<!--    const updateFile = (content: any, path: string, sha: string) => {-->
<!--      return new Promise(async (resolve) => {-->
<!--        const token = localStorage.token;-->
<!--        const userInfo = JSON.parse(localStorage.userInfo);-->
<!--        if (sha === "") {-->
<!--          // 新建文件-->
<!--          const response = await service.post("/repo/file", {-->
<!--            content: content,-->
<!--            login: userInfo.login,-->
<!--            path: path,-->
<!--            sha: sha,-->
<!--            token: token,-->
<!--          });-->
<!--          resolve(response);-->
<!--        } else {-->
<!--          // 更新文件-->
<!--          const response = await service.put("/repo/file", {-->
<!--            content: content,-->
<!--            login: userInfo.login,-->
<!--            path: path,-->
<!--            sha: sha,-->
<!--            token: token,-->
<!--          });-->
<!--          resolve(response);-->
<!--        }-->
<!--      });-->
<!--    };-->
<!--    init();-->

<!--    const selectedKeys = ref<string[]>([]);-->
<!--    const replaceFields = {-->
<!--      children: "child",-->
<!--      title: "name",-->
<!--      key: "path",-->
<!--    };-->
<!--    const onContextMenuClick = (node: any, menuKey: string) => {-->
<!--      console.log(`treeKey: ${node.path}, menuKey: ${menuKey}`);-->
<!--      //  新建文件-->
<!--      if (menuKey === "0") {-->
<!--        const path = node.path + "/新文件.md";-->
<!--        updateFile("", path, "").then((res: any) => {-->
<!--          let dataRef = node.dataRef;-->
<!--          dataRef.child.push({-->
<!--            name: res.data.name,-->
<!--            path: res.data.path,-->
<!--            isDir: false,-->
<!--            isShare: false,-->
<!--            sha: res.data.sha,-->
<!--            size: res.data.size,-->
<!--            url: res.data.url,-->
<!--            type: res.data.type,-->
<!--            html_url: res.data.html_url,-->
<!--            download_url: res.data.download_url,-->
<!--            createTime: getNowDate(),-->
<!--            updateTime: getNowDate(),-->
<!--          });-->
<!--          data.workspace.note = data.workspaceNotes[0].child;-->
<!--          updateWorkspace(data.workspace, data.workspaceSha);-->
<!--        });-->
<!--      }-->

<!--      if (menuKey === "1") {-->
<!--        node.child.unshift({-->
<!--          name: "新文件夹",-->
<!--          path: `${node.path}/新文件夹`,-->
<!--          isDir: true,-->
<!--          isShare: false,-->
<!--          sha: "",-->
<!--          size: 0,-->
<!--          url: "",-->
<!--          type: "",-->
<!--          html_url: "",-->
<!--          download_url: "",-->
<!--          createTime: getNowDate(),-->
<!--          updateTime: getNowDate(),-->
<!--          child: [],-->
<!--        });-->
<!--        updateWorkspace(data.workspace, data.workspaceSha);-->
<!--      }-->
<!--      if (menuKey === "2") {-->
<!--        // {-->
<!--        //   "name": "新文件夹",-->
<!--        //     "path": `${treeKey}/新文件夹`,-->
<!--        //     "updateTime": getNowDate(),-->
<!--        // }-->
<!--      }-->
<!--      //  删除-->
<!--      if (menuKey === "3") {-->
<!--        if (node.isDir && node.child.length) {-->
<!--          return message.error("存在文件，不允许删除");-->
<!--        }-->

<!--        let arr: any = {-->
<!--          note: [],-->
<!--          trash: [],-->
<!--          share: [],-->
<!--        };-->
<!--        updateWorkspace(arr, data.workspaceSha);-->
<!--      }-->
<!--    };-->
<!--    const updateName = (node) => {};-->

<!--    //  点击树节点触发-->
<!--    function selectTreeNode(selectedKeys, { node }) {-->
<!--      data.activeTreeNode = node.dataRef;-->
<!--      const { isDir, path } = node.dataRef;-->
<!--      if (!isDir) {-->
<!--        epoContent(path).then((res: any) => {-->
<!--          data.activeTreeNode.name = res.name;-->
<!--          data.activeTreeNode.sha = res.sha;-->
<!--          data.vditor.setValue(res.content);-->
<!--        });-->
<!--      }-->
<!--    }-->

<!--    //  创建笔记的文件路径信息-->
<!--    function createNote() {-->
<!--      const dataRef = data.activeTreeNode[0].dataRef;-->
<!--      let activeNote = {-->
<!--        name: "新文件.md",-->
<!--        path: `${dataRef.path}/新文件.md`,-->
<!--        isDir: false,-->
<!--        isShare: false,-->
<!--        sha: "",-->
<!--        size: 0,-->
<!--        url: "",-->
<!--        type: "",-->
<!--        html_url: "",-->
<!--        download_url: "",-->
<!--        createTime: getNowDate(),-->
<!--        updateTime: getNowDate(),-->
<!--      };-->
<!--      data.activeNote = activeNote;-->
<!--      data.activeTreeNode[0].dataRef.child.push(activeNote);-->

<!--      updateWorkspace(data.workspace, data.workspaceSha);-->

<!--      updateFile("", activeNote.path, "");-->
<!--    }-->

<!--    //  新建&&更新文件-->
<!--    function updateNote() {-->
<!--      const content = data.vditor.getValue(),-->
<!--        path = data.activeTreeNode.path,-->
<!--        sha = data.activeTreeNode.sha;-->
<!--      updateFile(content, path, sha).then((res: any) => {-->
<!--        data.activeTreeNode = {-->
<!--          name: res.data.name,-->
<!--          path: res.data.path,-->
<!--          isDir: false,-->
<!--          isShare: false,-->
<!--          sha: res.data.sha,-->
<!--          size: res.data.size,-->
<!--          url: res.data.url,-->
<!--          type: res.data.type,-->
<!--          html_url: res.data.html_url,-->
<!--          download_url: res.data.download_url,-->
<!--          createTime: getNowDate(),-->
<!--          updateTime: getNowDate(),-->
<!--          child: [],-->
<!--        };-->
<!--      });-->
<!--    }-->

<!--    return {-->
<!--      ...toRefs(data),-->
<!--      selectedKeys,-->
<!--      replaceFields,-->
<!--      selectTreeNode,-->
<!--      onContextMenuClick,-->
<!--      createNote,-->
<!--      updateNote,-->
<!--      updateName,-->
<!--    };-->
<!--  },-->
<!--});-->
<!--</script>-->

<!--<style scoped lang="scss">-->
<!--@import url(vditor/dist/index.css);-->

<!--.custom-resizer {-->
<!--  width: 100%;-->
<!--  height: 100vh;-->
<!--}-->

<!--.custom-resizer > .pane {-->
<!--  text-align: left;-->
<!--  padding: 15px;-->
<!--  overflow: hidden;-->
<!--  background: #fff;-->
<!--  border: 1px solid #ccc;-->

<!--  &.menu {-->
<!--    width: 20%;-->
<!--    max-width: 250px;-->
<!--  }-->

<!--  &.catalogue {-->
<!--    width: 300px;-->
<!--    background-color: #f4f6f9;-->
<!--  }-->
<!--}-->

<!--.custom-resizer > .multipane-resizer {-->
<!--  margin: 0;-->
<!--  left: 0;-->
<!--  position: relative;-->

<!--  &:before {-->
<!--    display: block;-->
<!--    content: "";-->
<!--    width: 3px;-->
<!--    height: 40px;-->
<!--    position: absolute;-->
<!--    top: 50%;-->
<!--    left: 50%;-->
<!--    margin-top: -20px;-->
<!--    margin-left: -1.5px;-->
<!--    border-left: 1px solid #ccc;-->
<!--    border-right: 1px solid #ccc;-->
<!--  }-->

<!--  &:hover {-->
<!--    &:before {-->
<!--      border-color: #999;-->
<!--    }-->
<!--  }-->
<!--}-->

<!--.menu {-->
<!--  .avatar {-->
<!--    display: flex;-->
<!--    justify-content: center;-->
<!--    align-items: center;-->
<!--    flex-direction: column;-->
<!--    margin-bottom: 20px;-->
<!--  }-->
<!--}-->

<!--.catalogue {-->
<!--  .ant-divider-horizontal {-->
<!--    margin: 15px 0;-->
<!--  }-->

<!--  .info {-->
<!--    display: flex;-->
<!--    align-items: center;-->
<!--    justify-content: space-between;-->

<!--    p {-->
<!--      margin: 0;-->
<!--    }-->
<!--  }-->
<!--}-->

<!--.doc {-->
<!--  &-title {-->
<!--    display: flex;-->
<!--    margin-bottom: 10px;-->
<!--  }-->
<!--}-->
<!--</style>-->
