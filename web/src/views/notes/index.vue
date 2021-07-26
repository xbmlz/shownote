<template>
  <multipane class="custom-resizer" layout="vertical">
    <div class="pane menu" @contextmenu="rightClick($event,{})" @click="handleMenu">
      <div class="avatar">
        <div>
          <el-avatar
              shape="square"
              size="large"
              :src="userInfo.avatar_url"
          ></el-avatar>
          <h2>{{ userInfo.name }}</h2>
        </div>
        <el-button type="primary" @click="createDir" size="mini" icon="el-icon-plus" circle></el-button>
      </div>
      <el-tree
          ref="noteTree"
          highlight-current
          node-key="path"
          :data="notes"
          :props="defaultProps"
          :default-expanded-keys="expandedKeys"
          empty-text="暂无数据"
          @node-expand="handleNodeExpand"
          @node-collapse="handleNodeCollapse"
          @node-click="handleNodeClick"
          @node-contextmenu="rightClick">
        <template #default="{ node, data }">
          <div>
            <i class="el-icon-folder-opened" v-if="data.isDir"></i>
            <i class="el-icon-document" v-else></i>
            <span style="margin-left: 4px">{{ data.name }}</span>
          </div>
        </template>
      </el-tree>
    </div>
    <multipane-resizer></multipane-resizer>
    <div class="pane doc" :style="{ flexGrow: 1 }">
      <div v-show="activeNote.name" v-loading="isLoading" style="width: 100%;">
        <div class="doc-title">
          <el-input v-model="activeNote.name" readonly size="mini"></el-input>
          <el-button @click="updateNote" :loading="saveLoading" size="mini" type="primary" icon="el-icon-s-promotion">
            保存
          </el-button>
        </div>
        <div id="vditor"></div>
      </div>
      <el-empty v-show="!activeNote.name" description="暂无数据"></el-empty>
    </div>
  </multipane>

  <!-- 模态-文件及文件夹命名 -->
  <el-dialog v-model="isDialog" :title="dialogTitle" width="450px">
    <el-input v-model="rename" ref="rename" @keyup.enter="updateName"></el-input>
    <template #footer>
      <el-button size="small" @click="isDialog = false">取消</el-button>
      <el-button type="primary" size="small" @click="updateName">确定</el-button>
    </template>
  </el-dialog>

  <!--鼠标右键菜单栏 -->
  <div v-show="showRightMenu" class="menu-modal" oncontextmenu="return false;">
    <ul id="menu" class="right-menu">
      <li class="menu-item" @click="createFile" v-if="activeTreeNode.isDir">新建文件</li>
      <li class="menu-item" @click="createDir" v-if="activeTreeNode.isDir !== false">新建文件夹</li>
      <li class="menu-item" @click="updateDir" v-if="activeTreeNode.isDir">重命名</li>
      <li class="menu-item" @click="updateFileName" v-if="activeTreeNode.isDir === false">重命名</li>
      <li class="menu-item" @click="deleteFile" v-if="activeTreeNode.isDir === false">删除</li>
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
      saveLoading: false,
      isLoading: true, // 文件内容获取中
      isHideToolbar: false,
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
    async init() {
      const info = await service.get(`/user/info?token=${this.token}`);
      this.userInfo = info.data;
      this.initVditor()
      this.initData()
    },
    // 初始化编辑器
    initVditor() {
      this.vditor = new Vditor("vditor", {
        height: "100%",
        toolbarConfig: {
          pin: true,
          hide: this.isHideToolbar
        },
        upload: {
          url: 'http://10.0.2.172:8000/repo/upload',
          multiple: false,
          fieldName: 'file',
          extraData: {
            login: this.userInfo.login,
            token: localStorage.token
          }
        },
        cache: {
          enable: false,
        },
      });
    },
    //  初始化数据
    async initData() {
      this.initRepo().then(res => {
        let workspace = JSON.parse(res.content);
        const {notes, trash, share} = workspace;
        this.notes = notes;
        this.trash = trash;
        this.share = share;
        this.workspaceSha = res.sha;
      });
    },
    handleMenu(e) {
      this.$refs['noteTree'].setCurrentKey()
    },
    //  新建文件夹
    createDir() {
      this.isDialog = true;
      this.dialogTitle = '新建文件夹';
      this.type = 'addDir';
      this.$nextTick(() => {
        this.rename = '';
        this.$refs['rename'].focus()
      })
    },
    // 更新文件夹
    updateDir() {
      this.isDialog = true;
      this.dialogTitle = '修改文件夹';
      this.type = 'updateDir';
      this.$nextTick(() => {
        this.rename = this.activeTreeNode.name;
        this.$refs['rename'].focus()
      })
    },
    //  新建文件
    createFile() {
      this.isDialog = true;
      this.dialogTitle = '新建文件';
      this.type = 'addFile';
      this.$nextTick(() => {
        this.rename = '';
        this.$refs['rename'].focus()
      })
    },
    // 更新文件名
    updateFileName() {
      this.isDialog = true;
      this.dialogTitle = '修改文件';
      this.type = 'updateDir';
      this.$nextTick(() => {
        this.rename = this.activeTreeNode.name;
        this.$refs['rename'].focus()
      })
    },
    //  更新文件及目录结构
    updateNote() {
      if (!this.activeNote.name) {
        return this.$message.error('文件名称不能为空');
      }
      let params = {
        uid: this.activeNote.uid,
        content: this.vditor.getValue() || ' ',
        sha: this.activeNote.sha
      }
      this.saveLoading = true;
      this.updateFile(params).then(res => {
        this.$message.success(res.msg)
        this.saveLoading = false;
        this.activeNote = {
          name: this.activeNote.name,
          path: res.data.path,
          uid: res.data.uid,
          isDir: false,
          isShare: false,
          sha: res.data.sha,
          size: res.data.size,
          url: res.data.url,
          type: res.data.type,
          html_url: res.data.html_url,
          download_url: res.data.download_url,
          createTime: getNowDate(),
          updateTime: getNowDate(),
          child: [],
        };
        this.replaceTreeNode(this.activeNote, this.notes)
      }).catch(() => {
        this.saveLoading = false;
      });
    },
    //  更新目录空间节点信息
    replaceTreeNode(data, notes) {
      for (let i in notes) {
        let {path, child} = notes[i];
        if (path === data.path) {
          notes[i] = data;
          this.updateWorkspace()
        } else {
          this.replaceTreeNode(data, child)
        }
      }
    },
    //  移除目录空间节点信息
    removeTreeNode(data, notes) {
      for (let i in notes) {
        let {path, child} = notes[i];
        if (path === data.path) {
          notes.splice(i, 1);
          this.activeTreeNode = {};
          this.activeNote = {};
          this.updateWorkspace()
        } else {
          this.removeTreeNode(data, child)
        }
      }
    },
    /****************** 模态-新建修改文件夹 ********************/
    // 功能：1.新建文件夹名;2.修改文件夹名;3.新建文件名;4.修改文件名
    updateName() {
      if (this.type === 'addDir') {
        // 检测重命名
        let node = this.$refs['noteTree'].getNode(this.activeTreeNode.path);
        let childNode = node ? node.data.child : [];
        this.rename = this.validateName(this.rename, true, childNode);
        this.addDir()
      }
      if (this.type === 'addFile') {
        // 检测重命名
        let node = this.$refs['noteTree'].getNode(this.activeTreeNode.path);
        this.rename = this.validateName(this.rename, false, node.data.child);

        this.addFile()
      }
      if (this.type === 'updateDir' || this.type === 'updateFile') {
        // 检测重命名
        let node = this.$refs['noteTree'].getNode(this.activeTreeNode.path)
        let fileArr = [], parentData = node.parent.data
        if (parentData.constructor === Array) {
          fileArr = parentData;
        } else {
          fileArr = parentData.child
        }

        this.rename = this.validateName(this.rename, false, fileArr, node.path);
        // 更新
        this.activeTreeNode.name = this.rename;
        this.replaceTreeNode(this.activeTreeNode, this.notes)
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
          uid: '',
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
        path = `notes/${this.rename}`
        this.notes.push({
          name: this.rename,
          path: path,
          uid: '',
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
          name: this.rename,
          path: res.data.path,
          uid: res.data.uid,
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
          this.$nextTick(() => {
            this.$refs['noteTree'].setCurrentKey(res.data.path)
            //  创建成功后打开文件
            const currebtData = this.$refs['noteTree'].getCurrentNode()
            this.handleNodeClick(currebtData)
          })
        })
      })
    },
    /****************** 目录 ********************/
    // 保存当前展开的节点
    handleNodeExpand(data) {
      let flag = false
      this.expandedKeys.some(item => {
        if (item === data.path) { // 判断当前节点是否存在， 存在不做处理
          flag = true
          return true
        }
      })
      if (!flag) { // 不存在则存到数组里
        this.expandedKeys.push(data.path)
      }
    },
    // 删除关闭的节点
    handleNodeCollapse(data) {
      this.expandedKeys.some((item, i) => {
        if (item === data.path) {
          // 删除关闭节点
          this.expandedKeys.length = i
        }
      })
    },

    handleNodeClick(data) {
      this.activeTreeNode = data;
      if (!data.isDir) {
        this.activeNote = {...data};
        this.isLoading = true;
        this.getRepoFile(data.uid).then(res => {
          this.isLoading = false;
          this.activeTreeNode.sha = res.sha;
          this.vditor.focus();
          this.vditor.setValue(res.content);
        }).catch(() => {
          this.isLoading = false;
        });
      }
    },
    // 右键菜单
    rightClick(event, data) {
      if (JSON.stringify(data) === '{}') {
        this.$refs['noteTree'].setCurrentKey()
        this.vditor.disabled()
      }
      this.activeTreeNode = data;
      this.showRightMenu = false // 先把模态框关死，目的是：第二次或者第n次右键鼠标的时候 它默认的是true
      this.showRightMenu = true
      let menu = document.querySelector('#menu')
      menu.style.left = event.clientX + 10 + 'px'
      menu.style.top = event.clientY - 5 + 'px'
      // 给整个document添加监听鼠标事件，点击任何位置执行closeRightMenu方法，及时将菜单关闭
      document.addEventListener('click', this.closeRightMenu)
      event.preventDefault()
    },
    closeRightMenu() {
      this.showRightMenu = false
      // 及时关掉鼠标监听事件
      document.removeEventListener('click', this.closeRightMenu)
    },

    /**
     * 验证重名
     * @params name 名称
     * @params isDir 是否为文件夹
     * @params fileArr 所有同级目录
     * @params path 路径，用于修改时判断是否为同一文件，非必填
     **/
    validateName(name, isDir, fileArr, path) {
      const nameFun = (name = '') => {
        let fileName = name;
        let resultName = fileName.split('-');
        if (resultName.length > 1) {
          const lastVal = resultName[resultName.length - 1];
          const regPos = /^\d+(\.\d+)?$/; // 非负浮点数
          if (regPos.test(lastVal)) {
            resultName[resultName.length - 1] = parseInt(lastVal) + 1;
            name = resultName.join('-')
          } else {
            name = fileName + '-1'
          }
        } else {
          name = fileName + '-1'
        }
        return name;
      }

      // 默认文件名为新文件
      let result = fileArr.filter(item => item.name === name && item.isDir === isDir && item.path !== path);
      if (!result.length) {
        console.log(name, '最终命名')
        return name;
      } else {
        let data = result[0];
        name = nameFun(data.name)
        return this.validateName(name, isDir, fileArr, path)
      }
    },
    /****************** API ********************/
    // 获取仓库信息,初始化使用
    initRepo() {
      return new Promise(async (resolve) => {
        const res = await service.get("/repo/info", {
          params: {
            token: localStorage.token,
            login: this.userInfo.login,
          },
        });
        resolve(res.data);
      });
    },
    // 获取工作空间
    getWorkspace() {
      return new Promise(async (resolve) => {
        let params = {
          token: localStorage.token,
          login: this.userInfo.login,
        };
        service.get('/repo/workspace', params).then(res => {
          resolve(res.data);
        })
      })
    },
    //  更新目录JSON
    updateWorkspace() {
      return new Promise(async resolve => {
        const token = localStorage.token;
        const userInfo = JSON.parse(localStorage.userInfo);
        let content = {
          notes: this.notes,
          trash: this.trash,
          share: this.share,
        }
        const res = await service.put("/repo/workspace", {
          content: JSON.stringify(content),
          login: userInfo.login,
          sha: this.workspaceSha,
          token: token,
          uid: ''
        });
        let workspace = JSON.parse(res.data.content);
        const {notes, trash, share} = workspace;
        this.notes = notes;
        this.trash = trash;
        this.share = share;
        this.workspaceSha = res.data.sha;
        this.$nextTick(() => {
          if (this.activeTreeNode.path) {
            this.$refs['noteTree'].setCurrentKey(this.activeTreeNode.path);
          }
        })

        this.isDialog = false;
        resolve()
      })
    },
    // 获取内容
    getRepoFile(uid) {
      return new Promise(async (resolve) => {
        let params = {
          token: localStorage.token,
          login: this.userInfo.login,
          uid: uid,
        }
        const res = await service.get("/repo/file", {params});
        resolve(res.data);
      });
    },
    // 更新文件
    updateFile(data) {
      return new Promise(async resolve => {
        const {content, path, sha, uid} = data;
        let response, params = {
          content: content || ' ',
          login: this.userInfo.login,
          token: localStorage.token,
          path,
          uid,
          sha,
        }
        if (sha === "") {
          // 新建文件
          response = await service.post("/repo/file", params);
        } else {
          // 更新文件
          response = await service.put("/repo/file", params);
        }
        if (response.code === 0) {
          resolve(response);
        }
      });
    },
    // 删除文件
    deleteFile() {
      return new Promise(resolve => {
        const {uid, sha} = this.activeTreeNode;
        let params = {
          token: localStorage.token,
          login: this.userInfo.login,
          uid,
          sha
        }
        service.delete("/repo/file", {params}).then(res => {
          this.removeTreeNode(this.activeTreeNode, this.notes)
          this.$message.success(res.msg);
          resolve()
        });
      })
    }
  }
});
</script>

<style scoped lang="scss">
@import '@/styles/notes.scss';
</style>
