<template>
  <transition name="header">
    <div class="header-panel">
      <a-button
        class="setting-btn"
        @click="
          (e) => {
            drawerVisible = true;
          }
        "
      >
        <template #icon>
          <CaretDownFilled />
        </template>
        <!-- {{ipAddress}} -->
      </a-button>
    </div>
  </transition>
  <a-drawer
    placement="top"
    :closable="true"
    :height="120"
    v-model:visible="drawerVisible"
  >
    <guac-client-setting
      ref="setting"
      :client="client"
      :display="display"
      :mouse="mouse"
      :keyboard="keyboard"
      @windowScale="windowScale"
      @inputModeChange="inputModeChange"
      @mouseModeChange="mouseModeChange"
    ></guac-client-setting>
  </a-drawer>
  <!-- <div ref="viewport" id="viewport" style="position: relative"> -->
  <div id="display" class="display software-cursor"></div>
  <div class="client-bottom">
    <!-- <div id="nestedInputText" class="footer-panel">
      <a-input v-model:value="inputTextUse" class="target" id="inputTextUse" />
    </div> -->
    <!-- <transition name="footer" v-if="inputText"> -->
    <!-- <div id ="nestedInputText" class="footer-panel">
        <a-input
          v-model:value="inputTextUse"
          class="target"
          id="inputTextUse"
        />
        <a-button danger>Ctrl</a-button>
        <a-button danger>Alt</a-button>
        <a-button danger>Esc</a-button>
        <a-button danger>Tab</a-button>
      </div> -->
    <!-- </transition> -->
    <!-- <transition name="footer" v-if="iuputOsk">
      <div class="footer-panel">
        <a-input v-model:value="vbvb" class="target1" />
      </div>
    </transition> -->
  </div>
  <!-- </div> -->

  <div class="modal" v-if="alertShow">
    <!-- <h2>{{ title[status] }}</h2>
    <p>{{ message ? message : text[status] }}</p>
    <span class="rct" @click="$emit('reconnect')" v-if="canReconnect">
      Reconnect
    </span> -->

    <a-result
      :status="resultStatus"
      :title="title[status]"
      :sub-title="text[status]"
    >
      <template #extra v-if="resultStatus == 'error'">
        <a-button v-if="ableReconnect" type="primary" @click="connect()">
          재접속
        </a-button>
      </template>
    </a-result>
  </div>
  <!-- <guac-client-modal ref="modal" @reconnect="connect()" /> -->
</template>

<script>
import { ref } from "vue";
import axios from "axios";
import Guacamole from "guacamole-common-js";
import onScreenKeyboardLayout from "@/keyboard-layouts/en-us-qwerty.json";
import encrypt from "@/lib/encrypt";
import dis from "@/lib/display";
import states from "@/lib/states";
import clipboard from "@/lib/clipboard";
import GuacClientSetting from "./GuacClientSetting";
const hostname =
  process.env.VUE_APP_API_URL == ""
    ? window.location.hostname
    : process.env.VUE_APP_API_URL;
const wsUrl = "ws://" + hostname + ":8088/";
export default {
  components: {
    GuacClientSetting,
  },
  setup() {},
  props: {},
  data() {
    return {
      cryptKey: "IgmTQVMISq9t4Bj7iRz7kZklqzfoXuq1",
      client: ref(null),
      keyboard: ref(null),
      mouse: ref(null),
      localCursor: ref(true),
      scale: ref(0),
      connectionState: ref(states.IDLE),
      errorMessage: ref(""),
      inputText: ref(false),
      inputOsk: ref(false),
      inputTextUse: ref(""),
      drawerVisible: ref(false),
      sentText: ref([]),
      arguments: {},
      token: {
        connection: {
          type: "rdp",
          settings: { "ignore-cert": true },
        },
      },
      // ipAddress: hostname,
      alertShow: ref(true),
      resultStatus: ref(null),
      status: ref(null),
      title: {
        CONNECTING: "데스크톱 연결 중",
        DISCONNECTING: "연결 끊김",
        DISCONNECTED: "연결 끊김",
        UNSTABLE: "불안정",
        WAITING: "대기중...",
        CLIENT_ERROR: "클라이언트 오류",
        TUNNEL_ERROR: "WebSocket Tunneling 오류",
      },
      text: {
        CONNECTING: "데스크톱에 연결을 위해 응답을 기다리는 중...",
        DISCONNECTING: "데스크톱에 접속이 끊어졌습니다.",
        DISCONNECTED: "데스크톱에 접속이 끊어졌습니다.",
        UNSTABLE: "서버에 대한 네트워크 연결이 불안정합니다.",
        WAITING: "서버에 응답을 기다리는 중입니다.",
        CLIENT_ERROR: "서버에 오류가 발생했습니다.",
        TUNNEL_ERROR: "WebSocket Tunneling 중 오류가 발생하였습니다.",
      },
    };
  },
  created() {},
  mounted() {
    // this.connectionState = states.CONNECTING;
    this.connect();
  },
  computed: {
    ableReconnect() {
      return ["DISCONNECTED", "CLIENT_ERROR"].includes(this.status);
    },
  },
  watch: {
    connectionState(state) {
      this.alertShow = true;
      console.log("connectionState ", state);
      this.alertShowHandle(state, this.errorMessage);
    },
    childScale(val) {
      //console.log("childScale ::::::::: ", val);
    },
    client(val) {
      //console.log("client ::::::::: ", val);
    },
    keyboard(val) {
      //console.log("keyboard ::::::::: ", val);
    },
    mouse(val) {
      //console.log("mouse ::::::::: ", val);
    },
    inputTextUse(val) {
      //console.log("inputTextUse ::::::::: ", val);
      this.target = document.getElementById("inputTextUse");

      console.log("target.selectionStart", this.target.selectionStart);
      if (this.composingText) return;

      var i;
      var content = val;
      var TEXT_INPUT_PADDING = 4;
      var expectedLength = TEXT_INPUT_PADDING * 2;

      // If content removed, update
      if (content.length < 2) {
        // Calculate number of backspaces and send
        var backspaceCount = TEXT_INPUT_PADDING - this.target.selectionStart;
        for (i = 0; i < backspaceCount; i++) this.sendKeysym(0xff08);

        // Calculate number of deletes and send
        var deleteCount = expectedLength - content.length - backspaceCount;
        for (i = 0; i < deleteCount; i++) this.sendKeysym(0xffff);
      } else {
        this.sendString(content);
      }

      this.resetTextInputTarget(TEXT_INPUT_PADDING);
      //e.preventDefault();

      // 한글 입력인경우
      this.target.addEventListener("compositionstart", (event) => {
        // console.log(`generated characters were: ${event}`, event);
        this.composingText = true;
      });
      this.target.addEventListener("compositionend", (event) => {
        console.log("한글 입력 :: ", event);
        this.composingText = false;
      });
    },
  },
  methods: {
    connect() {
      let tunnel = new Guacamole.WebSocketTunnel(wsUrl);

      tunnel.onerror = (status) => {
        //this.connectionState = states.TUNNEL_ERROR;
      };
      // tunnel.onstatechange = (state) => {
      //   console.log("tunnel", state);
      //   switch (state) {
      //     case 0:
      //       this.connectionState = states.CONNECTING;
      //       break;
      //     case 1:
      //       this.connectionState = states.CONNECTED;
      //       break;
      //     case 2:
      //       this.connectionState = states.DISCONNECTED;
      //       break;
      //     case 3:
      //       this.connectionState = states.UNSTABLE;
      //       break;
      //   }
      // };
      this.client = new Guacamole.Client(tunnel);
      this.client.onerror = (error) => {
        this.client.disconnect();
        this.errorMessage = error.message;
        this.connectionState = states.CLIENT_ERROR;
      };
      this.client.onstatechange = (state) => {
        // console.log("client", state);
        switch (state) {
          case 0:
            this.connectionState = states.IDLE;
            break;
          case 1:
            this.connectionState = states.CONNECTING;
            break;
          case 2:
            this.connectionState = states.WAITING;
            break;
          case 3:
            this.connectionState = states.CONNECTED;
            break;
          case 4:
            this.connectionState = states.DISCONNECTING;
            break;
          case 5:
            this.connectionState = states.DISCONNECTED;
            break;
        }
      };
      this.client.onsync = () => {};

      this.display = this.client.getDisplay();
      this.appEl = document.getElementById("app");
      this.displayEl = document.getElementById("display");

      if (this.displayEl !== null) this.displayEl.innerHTML = "";
      this.displayEl.appendChild(this.display.getElement());

      //display 현재 창 크기로 적용
      const browerSize = dis();
      this.token.connection.settings.width = browerSize[0];
      this.token.connection.settings.height = browerSize[1];
      this.token.connection.settings.dpi = browerSize[2];

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
      if (
        Math.floor(Date.now() / 1000) -
          parseInt(this.token.connection.settings.timestamp) >
        100000000
        //10초로 변경 예정
      ) {
        this.connectionState = states.DISCONNECTED;
      } else {
        this.client.connect("token=" + encrypt(this.token));
      }

      // this.client.onargv = (stream, mimetype, name) => {
      //   if (mimetype !== "text/plain") return;
      //   const reader = new Guacamole.StringReader(stream);
      //   let value = "";
      //   reader.ontext = (text) => {
      //     value += text;
      //   };
      //   reader.onend = () => {
      //     const stream = this.client.createArgumentValueStream(
      //       "text/plain",
      //       name
      //     );
      //     stream.onack = (status) => {
      //       if (status.isError()) {
      //         return;
      //       }
      //       this.arguments[name] = value;
      //     };
      //   };
      // };

      // this.displayEl.addEventListener("contextmenu", (e) => {
      //   e.stopPropagation();
      //   if (e.preventDefault) {
      //     e.preventDefault();
      //   }
      //   e.returnValue = false;
      // });

      //데스크탑 마우스 설정
      this.keyboard = new Guacamole.Keyboard(document);
      this.keyboard.onkeydown = (keysym) => {
        this.isMenuShortcutPressed(keysym);
        if (this.isMenuShortcutPressed(keysym)) {
          this.keyboard.reset();
          setTimeout(() => {
            this.drawerVisible = true;
          }, 400);
        }
        this.client.sendKeyEvent(1, keysym);
      };
      this.keyboard.onkeyup = (keysym) => {
        this.client.sendKeyEvent(0, keysym);
      };

      this.touch = new Guacamole.Touch(this.displayEl);
      // console.log(this.touch);
      // this.touch.onEach(
      //   ["touchstart', 'touchmove', 'touchend"],
      //   this.handleTouchEvent
      // );

      // this.touch.touchstart = (event) => {
      //   //alert(111);
      // };

      this.touchScreen = new Guacamole.Mouse.Touchscreen(this.displayEl);
      this.touchScreen.onmousedown = (state) => {
        // console.log(this.display.getWidth(), this.display.getHeight());
        // console.log(
        //   this.appEl.offsetWidth,
        //   this.appEl.offsetHeight
        // );
        //this.resize();
      };
      this.touchPad = new Guacamole.Mouse.Touchpad(this.displayEl);

      this.mouse = new Guacamole.Mouse(this.displayEl);
      this.mouse.onmouseout = (e) => {
        if (!this.display) return;
        this.display.showCursor(false);
      };

      this.mouse.on("mousedown", document.body.focus.bind(document.body));
      this.mouse.onEach(["mousedown", "mousemove", "mouseup"], (event) => {
        if (!this.client || !this.display) return;

        event.stopPropagation();
        event.preventDefault();

        // Send mouse state, show cursor if necessary
        this.display.showCursor(this.localCursor);
        this.client.sendMouseState(event.state, true);
      });
      this.mouse.onmousedown = (state) => {
        // console.log(this.display.getWidth(), this.display.getHeight());
        // console.log(this.appEl.offsetWidth,this.appEl.offsetHeight);
        this.resizeWindowEvent();
      };
      this.display.onresize = (x, y) => {
        //console.log("onresize event :::::::: ", x, y);
        //this.client.sendSize(x, y);
        //this.resize();
      };

      this.displayEl.onclick = () => {
        this.displayEl.focus();
      };
      // this.displayEl.onfocus = () => {
      //   this.displayEl.className = "focus";
      // };
      // this.displayEl.onblur = () => {
      //   this.displayEl.className = "";
      // };

      window.addEventListener("resize", this.resizeWindowEvent);
      this.mouseModeChange("touchscreen");
      this.inputModeChange("none");
    },
    isMenuShortcutPressed(keysym) {
      var SHIFT_KEYS = ["65505", "65506"],
        ALT_KEYS = ["65513", "65514", "65027", "65511", "65512"],
        CTRL_KEYS = ["65507", "65508"],
        //F12 = ["65481"],
        MENU_KEYS = [
          "65505",
          "65506",
          "65507",
          "65508",
          "65511",
          "65512",
          "65513",
          "65514",
          //"65481",
        ];
      // Ctrl+Alt+Shift has NOT been pressed if any key is currently held
      // down that isn't Ctrl, Alt, or Shift

      if (!MENU_KEYS.includes(keysym.toString())) return false;

      //console.log("::::::::::::::::::::::", Object.keys(this.keyboard.pressed));
      const arr = Object.keys(this.keyboard.pressed);

      return !!(
        SHIFT_KEYS.filter((x) => arr.includes(x)).length > 0 &&
        (ALT_KEYS.filter((x) => arr.includes(x)).length > 0 ||
          CTRL_KEYS.filter((x) => arr.includes(x)).length > 0)
      );
    },
    inputModeChange(inputMethod) {
      this.drawerVisible = false;
      // Remove current keyboard
      if (this.onScreenKeyboard) {
        this.$refs.display.removeChild(this.onScreenKeyboard.getElement());
        this.onScreenKeyboard = null;
      }
      if (document.getElementById("nestedInputText")) {
        document.getElementById("nestedInputText").remove();
      }

      if (inputMethod == "none") {
        this.inputText = false;
        this.inputOsk = false;
      } else if (inputMethod == "text") {
        this.inputText = true;
        this.inputOsk = false;

        let div = document.createElement("div");
        div.id = "nestedInputText";
        div.className = "target";
        div.innerHTML =
          "<a-input v-model:value='inputTextUse' class='target' id='inputTextUse' />" +
          "<a-button danger>Ctrl</a-button>" +
          "<a-button danger>Alt</a-button>" +
          "<a-button danger>Esc</a-button>" +
          "<a-button danger>Tab</a-button>";
        this.displayEl.appendChild(div);
      } else if (inputMethod == "osk") {
        this.inputOsk = true;
        this.inputText = false;

        var layout = onScreenKeyboardLayout;
        this.onScreenKeyboard = new Guacamole.OnScreenKeyboard(layout);
        //console.log(this.onScreenKeyboard.getElement());
        this.$refs.display.appendChild(this.onScreenKeyboard.getElement());

        this.onScreenKeyboard.onkeydown = (keysym) => {
          this.client.sendKeyEvent(1, keysym);
        };

        // Broadcast keydown for each key released
        this.onScreenKeyboard.onkeyup = (keysym) => {
          this.client.sendKeyEvent(0, keysym);
        };
      }
      setTimeout(() => {
        this.resizeWindowEvent();
      }, 1000);
    },
    mouseModeChange(absolute) {
      this.drawerVisible = false;
      // this.touch.offEach(
      //   ["touchstart', 'touchmove', 'touchend"],
      //   this.handleTouchEvent
      // );

      // this.touch.onEach(
      //   ["touchstart', 'touchmove', 'touchend"],
      //   this.handleTouchEvent
      // );

      this.touchScreen.offEach(
        ["mousedown", "mousemove", "mouseup"],
        this.handleEmulatedMouseEvent
      );

      this.touchPad.offEach(
        ["mousedown", "mousemove", "mouseup"],
        this.handleEmulatedMouseEvent
      );

      if (absolute == "touchscreen") {
        this.touchScreen.onEach(
          ["mousedown", "mousemove", "mouseup"],
          this.handleEmulatedMouseEvent
        );
        // console.log(this.touchScreen);
        this.touchScreen.mousedown = (event) => {
          // console.log(event);
          //this.resize();
        };
      } else {
        this.touchPad.onEach(
          ["mousedown", "mousemove", "mouseup"],
          this.handleEmulatedMouseEvent
        );
        // console.log(this.touchPad);
        this.touchPad.mousedown = (event) => {
          // console.log(event);
          //this.resize();
        };
      }
    },
    handleTouchEvent(event) {
      // console.log("handleTouchEvent", event);
      // Do not attempt to handle touch state changes if the client
      // or display are not yet available
      if (!this.client || !this.display) return;
      event.preventDefault();

      // Send touch state, hiding local cursor
      this.display.showCursor(false);
      this.client.sendTouchState(event.state, true);
    },
    handleEmulatedMouseEvent(event) {
      //console.log("handleEmulatedMouseEvent", event);
      if (!this.client || !this.display) return;

      event.stopPropagation();
      event.preventDefault();

      // Ensure software cursor is shown
      this.display.showCursor(true);

      // Send mouse state, ensure cursor is visible
      this.scrollToMouse(event.state);
      this.client.sendMouseState(event.state, true);
    },
    scrollToMouse(mouseState) {
      const main = this.$refs.display;
      if (!main || !main.offsetWidth) {
        return;
      }
      // Determine mouse position within view
      const mouse_view_x =
        mouseState.x + this.displayEl.offsetLeft - main.scrollLeft;
      const mouse_view_y =
        mouseState.y + this.displayEl.offsetTop - main.scrollTop;

      // Determine viewport dimensions
      const view_width = main.offsetWidth;
      const view_height = main.offsetHeight;

      // Determine scroll amounts based on mouse position relative to document
      let scroll_amount_x;
      if (mouse_view_x > view_width)
        scroll_amount_x = mouse_view_x - view_width;
      else if (mouse_view_x < 0) scroll_amount_x = mouse_view_x;
      else scroll_amount_x = 0;

      let scroll_amount_y;
      if (mouse_view_y > view_height)
        scroll_amount_y = mouse_view_y - view_height;
      else if (mouse_view_y < 0) scroll_amount_y = mouse_view_y;
      else scroll_amount_y = 0;

      // Scroll (if necessary) to keep mouse on screen.
      main.scrollLeft += scroll_amount_x;
      main.scrollTop += scroll_amount_y;
    },
    resizeWindowEvent() {
      const pixelDensity = window.devicePixelRatio || 1;

      const width = this.appEl.offsetWidth * pixelDensity;
      const height = this.appEl.offsetHeight * pixelDensity;

      this.client.sendSize(width, height);

      if (this.onScreenKeyboard) {
        this.onScreenKeyboard.resize(this.$refs.display.offsetWidth);
      }
    },

    // setWindowSize() {
    //   // console.log("resize ::: this.childScale", this.childScale);
    //   setTimeout(() => {
    //     if (this.childScale > 0) {
    //       console.log(this.childScale, this.display.getScale());
    //       this.display.scale(this.childScale * 0.01);
    //       this.windowScale(this.childScale);
    //     } else {
    //       const scale = Math.min(
    //         document.getElementById("viewport").offsetWidth /
    //           Math.max(this.display.getWidth(), 1),
    //         document.getElementById("viewport").offsetHeight /
    //           Math.max(this.display.getHeight(), 1)
    //       );

    //       this.defaultScale =
    //         Math.round(scale * 100) > 300 ? 100 : Math.round(scale * 100);
    //       //console.log(this.defaultScale);
    //     }
    //   }, 1000);
    // },
    windowScale(curScale, defaultScale) {
      this.display.scale(curScale * 0.01);

      //확대시 스크롤 생성 및 해제
      if (curScale > defaultScale) this.appEl.style.overflow = "auto";
      else this.appEl.style.overflow = "hidden";
    },
    sendKeysym(keysym) {
      // console.log("sendKeysym [keysym] :::::::::", keysym);
      this.client.sendKeyEvent(1, keysym);
      this.client.sendKeyEvent(0, keysym);
      // $rootScope.$broadcast("guacSyntheticKeydown", keysym);
      // $rootScope.$broadcast("guacSyntheticKeyup", keysym);
    },
    sendCodepoint(codepoint) {
      if (codepoint === 10) {
        this.sendKeysym(0xff0d);
        // this.releaseStickyKeys();
        return;
      }

      var keysym = this.keysymFromCodepoint(codepoint);
      if (keysym) {
        this.sendKeysym(keysym);
        // this.releaseStickyKeys();
      }
    },
    keysymFromCodepoint(codepoint) {
      // Keysyms for control characters
      if (codepoint <= 0x1f || (codepoint >= 0x7f && codepoint <= 0x9f))
        return 0xff00 | codepoint;

      // Keysyms for ASCII chars
      if (codepoint >= 0x0000 && codepoint <= 0x00ff) return codepoint;

      // Keysyms for Unicode
      if (codepoint >= 0x0100 && codepoint <= 0x10ffff)
        return 0x01000000 | codepoint;

      return null;
    },
    sendString(content) {
      var sentText = "";

      // Send each codepoint within the string
      for (var i = 0; i < content.length; i++) {
        var codepoint = content.charCodeAt(i);
        if (codepoint !== 0x200b) {
          sentText += String.fromCharCode(codepoint);
          this.sendCodepoint(codepoint);
        }
      }

      // // Display the text that was sent

      this.sentText.push(sentText);

      // // Remove text after one second
      setTimeout(() => {
        this.sentText.shift();
      }, 1000);
      // console.log(this.sentText);
    },
    resetTextInputTarget(padding) {
      var TEXT_INPUT_PADDING_CODEPOINT = 0x200b;
      var paddingChar = String.fromCharCode(TEXT_INPUT_PADDING_CODEPOINT);

      this.inputTextUse = new Array(padding * 2 + 1).join(paddingChar);
      this.target.setSelectionRange(padding, padding);
    },
    alertShowHandle(state, message) {
      // console.log(state, message);
      this.status = state;

      switch (state) {
        case states.CONNECTING:
          this.resultStatus = "info";
          break;
        case states.WAITING:
          this.resultStatus = "info";
          break;
        case states.CONNECTED:
          this.alertShow = false;
          this.status = null;
          this.resultStatus = null;
          break;
        case states.DISCONNECTING:
          this.resultStatus = "error";
          break;
        case states.DISCONNECTED:
          this.resultStatus = "error";
          break;
        case states.TUNNEL_ERROR:
          this.resultStatus = "error";
          break;
        case states.CLIENT_ERROR:
          this.resultStatus = "error";
          break;
      }
    },
    // send(cmd) {
    //   if (!this.client) {
    //     return;
    //   }
    //   for (const c of cmd.data) {
    //     this.client.sendKeyEvent(1, c.charCodeAt(0));
    //   }
    // },
    // copy(cmd) {
    //   if (!this.client) {
    //     return;
    //   }
    //   clipboard.cache = {
    //     type: "text/plain",
    //     data: cmd.data,
    //   };
    //   clipboard.setRemoteClipboard(this.client);
    // },
    // handleMouseState(mouseState) {
    //   const scaledMouseState = Object.assign({}, mouseState, {
    //     x: mouseState.x / this.display.getScale(),
    //     y: mouseState.y / this.display.getScale(),
    //   });
    //   this.client.sendMouseState(scaledMouseState);
    // },
  },
};
</script>

<style lang="css">
@import "../assets/css/client.css";
</style>
