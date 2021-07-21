<template>
  <multipane class="custom-resizer" layout="vertical">
    <div class="pane menu" oncontextmenu="return false;">
      <div class="avatar">
        <div>
          <el-avatar
              shape="square"
              size="large"
              :src="userInfo.avatar_url"
          ></el-avatar>
          <h2>{{ userInfo.name }}</h2>
        </div>
        <el-button type="primary" @click="createDir" size="mini" icon="el-icon-plus"></el-button>
      </div>
      <el-tree
          ref="noteTree"
          highlight-current
          node-key="path"
          :data="notes"
          :props="defaultProps"
          :default-expanded-keys="expandedKeys"
          @node-click="handleNodeClick"
          @node-contextmenu="rightClick">
        <template #default="{ node, data }">
          <div>
            <i class="el-icon-folder-opened" v-if="data.isDir"></i>
            <i class="el-icon-document" v-else></i>
            <span>{{ data.name }}</span>
          </div>
        </template>
      </el-tree>
    </div>
    <multipane-resizer></multipane-resizer>
    <div class="pane doc" :style="{ flexGrow: 1 }">
      <div v-show="activeNote">
        <div class="doc-title">
          <el-input v-model="activeNote.name"></el-input>
          <el-button @click="updateNote">保存</el-button>
        </div>
        <div id="vditor"></div>
      </div>
      <el-empty v-show="!activeNote"></el-empty>
    </div>
  </multipane>
  <el-dialog v-model="isDialog" :title="dialogTitle" width="450px">
    <el-input v-model="rename"></el-input>
    <template #footer>
      <el-button size="small" @click="isDialog = false">取消</el-button>
      <el-button type="primary" size="small" @click="updateName">确定</el-button>
    </template>
  </el-dialog>

  <!--鼠标右键菜单栏 -->
  <div v-show="showRightMenu" class="menu-modal" @contextmenu="showRightMenu = false;">
    <ul id="menu" class="right-menu">
      <li class="menu-item" @click="createDir" v-if="activeTreeNode.isDir">新建文件夹</li>
      <li class="menu-item" @click="createNote" v-if="activeTreeNode.isDir">新建文件</li>
      <li class="menu-item" @click="createNote" v-if="activeTreeNode.isDir">重命名</li>
      <li class="menu-item" @click="createNote" v-if="!activeTreeNode.isDir">重命名1</li>
      <li class="menu-item">删除</li>
    </ul>
  </div>
</template>

<script>
import {ref, defineComponent, onMounted, reactive, toRefs} from "vue";
import {useRouter, useRoute} from "vue-router";
import service from "@/utils/request";
import {Multipane, MultipaneResizer} from "@/components/Multipane";
import {getNowDate} from "@/utils/date";
import Vditor from "vditor";

export default defineComponent({
  name: "App",
  components: {
    Multipane,
    MultipaneResizer,
  },
  data() {
    return {
      token: '',
      // 编辑器实例
      vditor: null,
      // 用户信息
      userInfo: {},
      //  当前选中节点信息
      showRightMenu: false,
      activeTreeNode: {},
      // 当前选中的文档信息
      workspaceSha: '',
      activeNote: {},
      expandedKeys: [],
      // 文件目录
      notes: [],
      defaultProps: {
        children: 'child',
        label: 'name'
      },
      trash: [],
      share: [],
      // dialog
      isDialog: false,
      dialogTitle: '新建文件夹',
      rename: '',
    };
  },
  mounted() {
    const router = useRouter();
    const route = useRoute();
    this.token = route.query.token || "";
    if (!this.token) {
      localStorage.token = ''
      router.push("/login");
    } else {
      localStorage.token = this.token;
      this.init()
    }
  },
  methods: {
    init() {
      this.initVditor()
      this.initData()
    },
    // 初始化编辑器
    initVditor() {
      this.vditor = new Vditor("vditor", {
        height: "100%",
        toolbarConfig: {
          pin: true,
        },
        cache: {
          enable: false,
        },
      });
    },
    //  初始化数据
    async initData() {
      const info = await service.get(`/user/info?token=${this.token}`);
      this.userInfo = info.data;
      window.localStorage.setItem("userInfo", JSON.stringify(info.data));
      this.epoContent(".shownote/workspace.json").then(res => {
        let workspace = JSON.parse(res.content);
        const {note, trash, share} = workspace;
        console.log(note)
        this.notes = note;
        this.trash = trash;
        this.share = share;
        this.workspaceSha = res.sha;
      });
    },
    //  新建文件夹
    createDir() {
      this.isDialog = true;
      this.dialogTitle = '新建文件夹';
      this.type = 'addDir';
    },
    //  新建文件
    createNote() {
      this.isDialog = true;
      this.dialogTitle = '新建文件';
      this.type = 'addFile';
    },

    updateNote() {

    },

    /****************** 模态-新建修改文件夹 ********************/
    // 功能：1.新建文件夹名;2.修改文件夹名;3.新建文件名;4.修改文件名
    updateName() {
      if (this.type === 'addDir') {
        this.addDir()
      }
      if (this.type === 'addFile') {
        this.addFile()
      }
    },
    // 新建文件夹
    addDir() {
      if (!this.activeTreeNode.isDir && this.activeTreeNode.path) {
        return alert('文件不允许')
      }
      let path = ''
      if (this.activeTreeNode.isDir) {
        path = `${this.activeTreeNode.path}/${this.rename}`
        this.$refs['noteTree'].append({
          name: this.rename,
          path: path,
          isDir: true,
          isShare: false,
          sha: "",
          size: 0,
          url: "",
          type: "",
          html_url: "",
          download_url: "",
          createTime: getNowDate(),
          updateTime: getNowDate(),
          child: [],
        }, this.activeTreeNode.path)
      } else {
        path = `note/${this.rename}`
        this.notes.push({
          name: this.rename,
          path: path,
          isDir: true,
          isShare: false,
          sha: "",
          size: 0,
          url: "",
          type: "",
          html_url: "",
          download_url: "",
          createTime: getNowDate(),
          updateTime: getNowDate(),
          child: [],
        })
      }
      this.updateWorkspace().then(() => {
        this.isDialog = false;
        this.expandedKeys.push(path);
      })
    },
    // 新建文件
    addFile() {
      if (!this.activeTreeNode.isDir && this.activeTreeNode.path) {
        return alert('文件不允许')
      }
      let params = {
        content: ' ',
        sha: '',
        path: `${this.activeTreeNode.path}/${this.rename}.md`,
      }
      this.updateFile(params).then(res => {
        this.$refs['noteTree'].append({
          name: res.data.name,
          path: res.data.path,
          isDir: false,
          isShare: false,
          sha: res.data.sha,
          size: res.data.size,
          url: res.data.url,
          type: res.data.type,
          html_url: res.data.html_url,
          download_url: res.data.download_url,
          createTime: res.data.createTime,
          updateTime: res.data.updateTime,
        }, this.activeTreeNode.path);

        this.updateWorkspace().then(() => {
          this.isDialog = false;
          this.expandedKeys.push(path);
        })
      })
    },
    /****************** 目录 ********************/
    handleNodeClick(data) {
      this.activeTreeNode = data;
      if (!data.isDir) {
        this.activeNote = data;
      }
    },
    // 右键菜单
    rightClick(event, data, node, obj) {
      this.activeTreeNode = data;
      this.showRightMenu = false // 先把模态框关死，目的是：第二次或者第n次右键鼠标的时候 它默认的是true
      this.showRightMenu = true
      let menu = document.querySelector('#menu')
      menu.style.left = event.clientX + 10 + 'px'
      menu.style.top = event.clientY - 5 + 'px'
      // 给整个document添加监听鼠标事件，点击任何位置执行closeRightMenu方法，及时将菜单关闭
      document.addEventListener('click', this.closeRightMenu)
    },
    closeRightMenu() {
      this.showRightMenu = false
      // 及时关掉鼠标监听事件
      document.removeEventListener('click', this.closeRightMenu)
    },
    /****************** API ********************/
    epoContent(path) {
      return new Promise(async (resolve) => {
        const token = localStorage.token;
        const info = JSON.parse(localStorage.userInfo);
        const res = await service.get("/repo/content", {
          params: {
            token: token,
            login: info.login,
            path: path,
          },
        });
        resolve(res.data);
      });
    },
    //  更新目录JSON
    updateWorkspace() {
      return new Promise(async resolve => {
        const token = localStorage.token;
        const userInfo = JSON.parse(localStorage.userInfo);
        let content = {
          note: this.notes,
          trash: this.trash,
          share: this.share,
        }
        const res = await service.put("/repo/file", {
          content: JSON.stringify(content),
          login: userInfo.login,
          path: ".shownote/workspace.json",
          sha: this.workspaceSha,
          token: token,
        });
        console.log(res, 'res--------')
        resolve()
      })

      // data.workspaceSha = res.data.sha;
    },
    // 更新文件
    updateFile(data) {
      return new Promise(async resolve => {
        const {content, path, sha} = data;
        const userInfo = JSON.parse(localStorage.userInfo);
        let response, params = {
          content: content,
          login: userInfo.login,
          path: path,
          sha: sha,
          token: localStorage.token,
        }
        if (sha === "") {
          // 新建文件
          response = await service.post("/repo/file", params);
        } else {
          // 更新文件
          response = await service.put("/repo/file", params);
        }
        console.log(response, 'response')
        if (response.code === 0) {
          resolve(response);
        }
      });
    }
  }
});
</script>

<style scoped lang="scss">
@import '@/styles/notes.scss';
</style>
