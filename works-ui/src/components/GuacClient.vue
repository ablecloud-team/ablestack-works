<template>
  <a-drawer
    placement="top"
    :closable="true"
    :height="660"
    v-model:visible="visible"
    :after-visible-change="afterVisibleChange"
  >
      <a-row :gutter="16">
        <a-col></a-col>
        <a-col></a-col>
        <a-col></a-col>
      </a-row>
    <a-row :gutter="16">
      <a-col :span="8"
        ><h3>입력 방법</h3>
        <div class="content">
          <div class="choice">
            <label
              ><input
                id="ime-none"
                name="input-method"
                ng-change="closeMenu()"
                ng-model="menu.inputMethod"
                type="radio"
                value="none"
              />
              없음</label
            >
            <p class="caption">
              <label for="ime-none"
                >사용중인 입력 방법이 없습니다. 키보드 입력은 열결된 물리적
                키보드에서 받아들여집니다.</label
              >
            </p>
          </div>
          <div class="choice">
            <label
              ><input
                id="ime-text"
                name="input-method"
                ng-change="closeMenu()"
                ng-model="menu.inputMethod"
                type="radio"
                value="text"
              />텍스트 입력</label
            >
            <p class="caption">
              <label for="ime-text">
                텍스트 입력을 허용하고, 입력된 텍스트를 바탕으로 키보드 이벤트를
                에뮬레이트 합니다. 이것은 물리적 키보드가 없는 휴대폰과 같은
                장치에 필요합니다.</label
              ><div class="figure">
              <label for="ime-text"
                ><img src="../assets/tablet-keys.svg" alt="" style="width: 30%"
              /></label>
            </div>
            </p>
          </div>
          <div class="choice">
            <label
              ><input
                id="ime-osk"
                name="input-method"
                ng-change="closeMenu()"
                ng-model="menu.inputMethod"
                type="radio"
                value="osk"
              />
              화상 키보드</label
            >
            <p class="caption">
              <label for="ime-osk">
                내장된 데스크톱 화상 키보드의 입력을 표시하고 허용합니다. 화상
                키보드는 다른 방법으로는 불가능할 수 있는 키 조합을 입력할 수
                있습니다. (Ctrl-Alt-Del 등)</label
              >
            </p>
          </div>
        </div>
      </a-col>
      <a-col :span="8"
        ><h3>마우스 에뮬레이션 모드</h3>
        <div class="content">
          <p>터치와 관련하여 원격 마우스가 어떻게 동작하는지 결정합니다.</p>
          <!-- Touchscreen -->
          <div class="choice">
            <input
              name="mouse-mode"
              ng-change="closeMenu()"
              ng-model="menu.emulateAbsoluteMouse"
              type="radio"
              ng-value="true"
              checked="checked"
              id="absolute"
            />
            <div class="figure">
              <label for="absolute"
                ><img src="../assets/touchscreen.svg" alt="터치 스크린" style="width: 50%"
              /></label>
              <p class="caption">
                <label for="absolute"
                  >탭하여 클릭합니다. 클릭은 터치 위치에서 발생합니다.</label
                >
              </p>
            </div>
          </div>
          <!-- Touchpad -->
          <div class="choice">
            <input
              name="mouse-mode"
              ng-change="closeMenu()"
              ng-model="menu.emulateAbsoluteMouse"
              type="radio"
              ng-value="false"
              id="relative"
            />
            <div class="figure">
              <label for="relative"
                ><img src="../assets/touchpad.svg" alt="터치 패드" style="width: 50%"
              /></label>
              <p class="caption">
                <label for="relative"
                  >드래그하여 마우스 포인터를 움직이고 탭하여 클릭합니다. 클릭은
                  마우스 포인터 위치에서 발생합니다.</label
                >
              </p>
            </div>
          </div>
        </div>
        </a-col>
        <a-col :span="8"
        ><h3>디스플레이</h3>
        <div class="content">
          <div id="zoom-settings">
            <div class="client-zoom">
              <div class="client-zoom-editor">
                <div ng-click="zoomOut()" class="client-zoom-out">
                  <img src="../assets/zoom-out.svg" alt="-" />
                </div>
                <div class="client-zoom-state">
                  <input
                    type="number"
                    guac-zoom-ctrl
                    ng-model="client.clientProperties.scale"
                    ng-model-options="{ updateOn: 'blur submit' }"
                    ng-change="zoomSet()"
                  />%
                </div>
                <div ng-click="zoomIn()" class="client-zoom-in">
                  <img src="../assets/zoom-in.svg" alt="+" />
                </div>
              </div>
              <div class="client-zoom-autofit">
                <label
                  ><input
                    ng-model="client.clientProperties.autoFit"
                    ng-change="changeAutoFit()"
                    ng-disabled="autoFitDisabled()"
                    type="checkbox"
                    id="auto-fit"
                  />
                  브라우저 창에 자동으로 맞춤</label
                >
              </div>
            </div>
          </div>
        </div>
      </a-col>
    </a-row>
  </a-drawer>
  <div class="viewport" ref="viewport">
    <div ref="displayFrame" class="display" tabindex="0" />
    <modal ref="modal" @reconnect="connect()" />
    <div class="headerbar">
      <transition name="slide">
        <div class="header-panel">
          <a-button :size="small" @click="showDrawer">
            <template #icon><CaretDownFilled /> </template>
          </a-button>
        </div>
      </transition>
      <!-- <transition name="slide">
        <div v-if="guacMenu" class="sidebar-panel" id="guac-menu">
          <ul class="sidebar-panel-nav" id="clipboard-settings">
            <h3>클립보드</h3>
            <div class="content">
              <p>
                데스크톱에서 복사하거나 잘라낸 텍스트가 여기에 표시됩니다.
                텍스트 변경 사항은 원격 클립보드에 직접 적용됩니다.
              </p>
              <textarea class="clipboard"></textarea>
            </div>
          </ul>
          <ul class="sidebar-panel-nav" id="keyboard-settings">
            <h3>입력 방법</h3>
            <div class="content">
              <div class="choice">
                <label
                  ><input
                    id="ime-none"
                    name="input-method"
                    ng-change="closeMenu()"
                    ng-model="menu.inputMethod"
                    type="radio"
                    value="none"
                  />
                  없음</label
                >
                <p class="caption">
                  <label for="ime-none"
                    >사용중인 입력 방법이 없습니다. 키보드 입력은 열결된 물리적
                    키보드에서 받아들여집니다.</label
                  >
                </p>
              </div>
              <div class="choice">
                <div class="figure">
                  <label for="ime-text"
                    ><img src="../assets/tablet-keys.svg" alt=""
                  /></label>
                </div>
                <label
                  ><input
                    id="ime-text"
                    name="input-method"
                    ng-change="closeMenu()"
                    ng-model="menu.inputMethod"
                    type="radio"
                    value="text"
                  />
                  텍스트 입력</label
                >
                <p class="caption">
                  <label for="ime-text">
                    텍스트 입력을 허용하고, 입력된 텍스트를 바탕으로 키보드
                    이벤트를 에뮬레이트 합니다. 이것은 물리적 키보드가 없는
                    휴대폰과 같은 장치에 필요합니다.</label
                  >
                </p>
              </div>
              <div class="choice">
                <label
                  ><input
                    id="ime-osk"
                    name="input-method"
                    ng-change="closeMenu()"
                    ng-model="menu.inputMethod"
                    type="radio"
                    value="osk"
                  />
                  화상 키보드</label
                >
                <p class="caption">
                  <label for="ime-osk">
                    내장된 데스크톱 화상 키보드의 입력을 표시하고 허용합니다.
                    화상 키보드는 다른 방법으로는 불가능할 수 있는 키 조합을
                    입력할 수 있습니다. (Ctrl-Alt-Del 등)</label
                  >
                </p>
              </div>
            </div>
          </ul>
          <ul class="sidebar-panel-nav" id="mouse-settings">
            <h3>마우스 에뮬레이션 모드</h3>
            <div class="content">
              <p>터치와 관련하여 원격 마우스가 어떻게 동작하는지 결정합니다.</p>-->
              <!-- Touchscreen -->
              <!-- <div class="choice">
                <input
                  name="mouse-mode"
                  ng-change="closeMenu()"
                  ng-model="menu.emulateAbsoluteMouse"
                  type="radio"
                  ng-value="true"
                  checked="checked"
                  id="absolute"
                />
                <div class="figure">
                  <label for="absolute"
                    ><img src="../assets/touchscreen.svg" alt="터치 스크린"
                  /></label>
                  <p class="caption">
                    <label for="absolute"
                      >탭하여 클릭합니다. 클릭은 터치 위치에서
                      발생합니다.</label
                    >
                  </p>
                </div>
              </div> -->
              <!-- Touchpad -->
              <!-- <div class="choice">
                <input
                  name="mouse-mode"
                  ng-change="closeMenu()"
                  ng-model="menu.emulateAbsoluteMouse"
                  type="radio"
                  ng-value="false"
                  id="relative"
                />
                <div class="figure">
                  <label for="relative"
                    ><img src="../assets/touchpad.svg" alt="터치 패드"
                  /></label>
                  <p class="caption">
                    <label for="relative"
                      >드래그하여 마우스 포인터를 움직이고 탭하여 클릭합니다.
                      클릭은 마우스 포인터 위치에서 발생합니다.</label
                    >
                  </p>
                </div>
              </div>
            </div>
          </ul>
          <ul class="sidebar-panel-nav" id="display-settings">
            <h3>디스플레이</h3>
            <div class="content">
              <div id="zoom-settings">
                <div class="client-zoom">
                  <div class="client-zoom-editor">
                    <div ng-click="zoomOut()" class="client-zoom-out">
                      <img src="../assets/zoom-out.svg" alt="-" />
                    </div>
                    <div class="client-zoom-state">
                      <input
                        type="number"
                        guac-zoom-ctrl
                        ng-model="client.clientProperties.scale"
                        ng-model-options="{ updateOn: 'blur submit' }"
                        ng-change="zoomSet()"
                      />%
                    </div>
                    <div ng-click="zoomIn()" class="client-zoom-in">
                      <img src="../assets/zoom-in.svg" alt="+" />
                    </div>
                  </div>
                  <div class="client-zoom-autofit">
                    <label
                      ><input
                        ng-model="client.clientProperties.autoFit"
                        ng-change="changeAutoFit()"
                        ng-disabled="autoFitDisabled()"
                        type="checkbox"
                        id="auto-fit"
                      />
                      브라우저 창에 자동으로 맞춤</label
                    >
                  </div>
                </div>
              </div>
            </div>
          </ul>
        </div>
      </transition> --> 
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import Guacamole from "guacamole-common-js";
import encrypt from "@/lib/encrypt";
import dis from "@/lib/display";
import states from "@/lib/states";
import clipboard from "@/lib/clipboard";
import Modal from "./Modal";

const hostname =
  process.env.VUE_APP_API_URL == ""
    ? window.location.hostname
    : process.env.VUE_APP_API_URL;
const wsUrl = "ws://" + hostname + ":8088/";
export default {
  components: {
    Modal,
  },
  setup() {
    const visible = ref(false);

    const afterVisibleChange = (bool) => {
      console.log("visible", bool);
    };

    const showDrawer = () => {
      visible.value = true;
    };

    return {
      visible,
      afterVisibleChange,
      showDrawer,
    };
  },
  props: {},
  data() {
    return {
      cryptKey: "IgmTQVMISq9t4Bj7iRz7kZklqzfoXuq1",
      connected: false,
      guacMenu: false,
      displayFrame: null,
      client: null,
      keyboard: null,
      mouse: null,
      connectionState: states.IDLE,
      errorMessage: "",
      arguments: {},
      token: {
        connection: {
          type: "rdp",
          settings: { "ignore-cert": true },
        },
      },
    };
  },
  create() {
    new Guacamole.OnScreenKeyboard.Layout(template);
  },
  watch: {
    connectionState(state) {
      this.$refs.modal.show(state, this.errorMessage);
    },
  },
  methods: {
    closeSidebarPanel() {
      this.guacMenu = false;
    },
    send(cmd) {
      if (!this.client) {
        return;
      }
      for (const c of cmd.data) {
        this.client.sendKeyEvent(1, c.charCodeAt(0));
      }
    },
    copy(cmd) {
      if (!this.client) {
        return;
      }
      clipboard.cache = {
        type: "text/plain",
        data: cmd.data,
      };
      clipboard.setRemoteClipboard(this.client);
    },
    handleMouseState(mouseState) {
      const scaledMouseState = Object.assign({}, mouseState, {
        x: mouseState.x / this.displayFrame.getScale(),
        y: mouseState.y / this.displayFrame.getScale(),
      });
      this.client.sendMouseState(scaledMouseState);
    },
    resize() {
      const elm = this.$refs.viewport;
      if (!elm || !elm.offsetWidth) {
        return;
      }
      let pixelDensity = window.devicePixelRatio || 1;
      const width = elm.offsetWidth * pixelDensity;
      const height = elm.offsetHeight * pixelDensity;
      if (
        this.displayFrame.getWidth() !== width ||
        this.displayFrame.getHeight() !== height
      ) {
        this.client.sendSize(width, height);
      }
      setTimeout(() => {
        const scale = Math.min(
          window.innerWidth / Math.max(this.displayFrame.getWidth(), 1),
          window.innerHeight / Math.max(this.displayFrame.getHeight(), 1)
        );
        this.displayFrame.scale(scale);
      }, 100);
    },
    connect() {
      let tunnel = new Guacamole.WebSocketTunnel(wsUrl);
      // console.log(tunnel);
      //reconnect 경우 초기화
      if (this.client) {
        this.displayFrame.scale(0);
        this.keyboard.onkeydown = this.keyboard.onkeyup = () => {};
      }

      this.client = new Guacamole.Client(tunnel);
      console.log(this.client);
      clipboard.install(this.client);

      tunnel.onerror = (status) => {
        // console.log(status);
        this.connectionState = states.TUNNEL_ERROR;
      };

      tunnel.onstatechange = (state) => {
        console.log(state);
        switch (state) {
          case Guacamole.Tunnel.State.CONNECTING:
            this.connectionState = states.CONNECTING;
            break;
          case Guacamole.Tunnel.State.OPEN:
            this.connectionState = states.CONNECTED;
            break;
          case Guacamole.Tunnel.State.UNSTABLE:
            break;
          case Guacamole.Tunnel.State.CLOSED:
            this.connectionState = states.DISCONNECTED;
            break;
        }
      };

      this.client.onstatechange = (clientState) => {
        console.log(clientState);
        switch (clientState) {
          case 0:
            this.connectionState = states.IDLE;
            break;
          case 1:
            break;
          case 2:
            this.connectionState = states.WAITING;
            break;
          case 3:
            this.connectionState = states.CONNECTED;
            window.addEventListener("resize", this.resize);
            this.$refs.viewport.addEventListener("mouseenter", this.resize);
            clipboard.setRemoteClipboard(this.client);
            break;
          case 4:
            break;
          case 5:
            break;
        }
      };

      this.client.onerror = (error) => {
        this.client.disconnect();
        this.guacMenu = false;
        this.errorMessage = error.message;
        this.connectionState = states.CLIENT_ERROR;
      };
      this.client.onsync = () => {};

      this.client.onargv = (stream, mimetype, name) => {
        if (mimetype !== "text/plain") return;
        const reader = new Guacamole.StringReader(stream);
        let value = "";
        reader.ontext = (text) => {
          value += text;
        };
        reader.onend = () => {
          const stream = this.client.createArgumentValueStream(
            "text/plain",
            name
          );
          stream.onack = (status) => {
            if (status.isError()) {
              return;
            }
            this.arguments[name] = value;
          };
        };
      };

      this.client.onclipboard = clipboard.onClipboard;
      this.displayFrame = this.client.getDisplay();
      console.log(this.client.getDisplay().getHeight());
      //console.log(this.display);
      const displayElm = this.$refs.displayFrame;
      //console.log(displayElm);
      //console.log(this.display.getElement());
      displayElm.appendChild(this.displayFrame.getElement());
      displayElm.addEventListener("contextmenu", (e) => {
        e.stopPropagation();
        if (e.preventDefault) {
          e.preventDefault();
        }
        e.returnValue = false;
      });

      //display 현재 창 크기로 적용
      const dis_chg = dis();
      this.token.connection.settings.width = dis_chg[0];
      this.token.connection.settings.height = dis_chg[1];
      this.token.connection.settings.dpi = dis_chg[2];

      //파라미터로 넘어온 값 파라미터값 복호화
      const decrypted = this.$CryptoJS.AES.decrypt(
        atob(this.$route.query.enc),
        this.cryptKey
      ).toString(this.$CryptoJS.enc.Utf8);
      //console.log(JSON.parse(decrypted));

      //복호화 한 값 JSON 형식으로 변경 후 키,값 구분하여 token에 세팅
      const query = JSON.parse(decrypted);
      for (const key in query) {
        this.token.connection.settings[key] = query[key];
      }
      // token.connection.settings.port = this.$route.query.port;
      // token.connection.settings.hostname = this.$route.query.hostname;
      // token.connection.settings.domain = this.$route.query.domain;
      // token.connection.settings.password = this.$route.query.password;
      // token.connection.settings.username = this.$route.query.username;

      const encrypt_token = encrypt(this.token);
      this.client.connect("token=" + encrypt_token);

      //데스크탑 마우스 설정
      this.mouse = new Guacamole.Mouse(displayElm);
      // console.log(this.mouse);
      this.mouse.onmouseout = () => {
        if (!this.displayFrame) return;
        this.displayFrame.showCursor(false);
      };
      displayElm.onclick = () => {
        displayElm.focus();
      };
      displayElm.onfocus = () => {
        displayElm.className = "focus";
      };
      displayElm.onblur = () => {
        displayElm.className = "";
      };

      //데스크탑 키보드 설정
      var isShift = false;
      var isCtrl = false;
      var isAlt = false;

      this.keyboard = new Guacamole.Keyboard(document);
      this.keyboard.onkeydown = (keysym) => {
        if (keysym === 65505) isShift = true;
        if (keysym === 65512) isCtrl = true;
        if (keysym === 65507) isAlt = true;

        //Sidebar 설정 (shift+alt+ctrl)
        if (isCtrl == true && isShift == true && isAlt == true) {
          this.showDrawer();
          //this.guacMenu = !this.guacMenu;
        }
        this.client.sendKeyEvent(1, keysym);
      };
      this.keyboard.onkeyup = (keysym) => {
        if (keysym === 65505) isShift = false;
        if (keysym === 65512) isCtrl = false;
        if (keysym === 65507) isAlt = false;
        this.client.sendKeyEvent(0, keysym);
      };

      // console.log(this.mouse);
      this.mouse.onmousedown =
        this.mouse.onmouseup =
        this.mouse.onmousemove =
          this.handleMouseState;
      setTimeout(() => {
        this.resize();
        displayElm.focus();
      }, 1000);
    },
  },
  mounted() {
    this.connected = true;
    this.connect();
  },
};
</script>

<style scoped>
.display {
  overflow: hidden;
  width: 100%;
  height: 100%;
}
.viewport {
  position: relative;
  width: 100%;
  height: 100%;
}
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.2s ease;
}
.slide-enter,
.slide-leave-to {
  transform: translateX(-100%);
  transition: all 150ms ease-in 0s;
}
.header-panel {
  position: fixed;
  left: 0;
  top: 0;
  z-index: 999;
  text-align: center;
  width: 100%;
}

.sidebar-panel {
  overflow-y: auto;
  background-color: #fff;
  position: fixed;
  left: 0;
  top: 0;
  height: 100vh;
  z-index: 999;
  padding: 1rem 10px 1rem 10px;
  width: 465px;
}
ul.sidebar-panel-nav {
  list-style-type: none;
}
ul.sidebar-panel-nav > h3 {
  margin: 0;
  padding: 0;
  padding-bottom: 1em;
}
ul.sidebar-panel-nav ~ ul.sidebar-panel-nav h3 {
  padding-top: 1em;
}
.clipboard {
  position: relative;
  border: 1px solid #aaa;
  -moz-border-radius: 0.25em;
  -webkit-border-radius: 0.25em;
  -khtml-border-radius: 0.25em;
  border-radius: 0.25em;
  width: 100%;
  height: 2in;
  white-space: pre;
  font-size: 1em;
  overflow: auto;
  padding: 0.25em;
}
#guac-menu .content {
  padding: 0;
  margin: 0;
  /* IE10 */
  display: -ms-flexbox;
  -ms-flex-align: stretch;
  -ms-flex-direction: column;
  /* Ancient Mozilla */
  display: -moz-box;
  -moz-box-align: stretch;
  -moz-box-orient: vertical;
  /* Ancient WebKit */
  display: -webkit-box;
  -webkit-box-align: stretch;
  -webkit-box-orient: vertical;
  /* Old WebKit */
  display: -webkit-flex;
  -webkit-align-items: stretch;
  -webkit-flex-direction: column;
  /* W3C */
  display: flex;
  align-items: stretch;
  flex-direction: column;
}
#guac-menu .content > * {
  margin: 0;
  -ms-flex: 0 0 auto;
  -moz-box-flex: 0;
  -webkit-box-flex: 0;
  -webkit-flex: 0 0 auto;
  flex: 0 0 auto;
}
#guac-menu .content > * + * {
  margin-top: 1em;
}
#guac-menu .header h2 {
  white-space: nowrap;
  overflow: hidden;
  width: 100%;
  text-overflow: ellipsis;
}
#guac-menu #mouse-settings .choice {
  text-align: center;
}
#guac-menu #mouse-settings .choice .figure {
  display: inline-block;
  vertical-align: middle;
  width: 75%;
  max-width: 320px;
}
#guac-menu #keyboard-settings .caption {
  font-size: 0.9em;
  margin-left: 2em;
  margin-right: 2em;
}
#guac-menu #mouse-settings .figure .caption {
  text-align: center;
  font-size: 0.9em;
}
#guac-menu #mouse-settings .figure img {
  display: block;
  width: 100%;
  max-width: 320px;
  margin: 1em auto;
}
#guac-menu #keyboard-settings .figure {
  float: right;
  max-width: 30%;
  margin: 1em;
}
#guac-menu #keyboard-settings .figure img {
  width: 100%;
}
#guac-menu #zoom-settings {
  text-align: center;
}
.client-zoom .client-zoom-out,
.client-zoom .client-zoom-in,
.client-zoom .client-zoom-state {
  display: inline-block;
  vertical-align: middle;
}

.client-zoom .client-zoom-out,
.client-zoom .client-zoom-in {
  max-width: 3em;
  border: 1px solid rgba(0, 0, 0, 0.5);
  background: rgba(0, 0, 0, 0.1);
  border-radius: 2em;
  margin: 0.5em;
  cursor: pointer;
}

.client-zoom .client-zoom-out img,
.client-zoom .client-zoom-in img {
  width: 100%;
  opacity: 0.5;
}

.client-zoom .client-zoom-out:hover,
.client-zoom .client-zoom-in:hover {
  border: 1px solid rgba(0, 0, 0, 1);
  background: #cda;
}

.client-zoom .client-zoom-out:hover img,
.client-zoom .client-zoom-in:hover img {
  opacity: 1;
}

.client-zoom .client-zoom-state {
  font-size: 1.5em;
}

.client-zoom .client-zoom-autofit {
  text-align: left;
  margin-top: 1em;
}

.client-zoom .client-zoom-state input {
  width: 2em;
  font-size: 1em;
  padding: 0;
  background: transparent;
  border-color: rgba(0, 0, 0, 0.125);
}

.client-zoom .client-zoom-state input::-webkit-inner-spin-button,
.client-zoom .client-zoom-state input::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>
