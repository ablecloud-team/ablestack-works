<template>
  <div class="modal">
    <!-- <h2>{{ title[status] }}</h2>
    <p>{{ message ? message : text[status] }}</p>
    <span class="rct" @click="$emit('reconnect')" v-if="canReconnect">
      Reconnect
    </span> -->

    <a-result
      v-if="status"
      :status="title[status]"
      :title="message ? message : text[status]"
      sub-title="머여?"
    >
      <template #extra>
        <a-button
          v-if="canReconnect"
          type="primary"
          @click="$emit('reconnect')"
        >
          재접속
        </a-button>
      </template>
    </a-result>
  </div>
</template>
<script>
import states from "@/lib/states";
export default {
  data() {
    return {
      status: null,
      message: "",
      title: {
        CONNECTING: "Connecting",
        DISCONNECTED: "Disconnected",
        UNSTABLE: "Unstable",
        WAITING: "Waiting",
        CLIENT_ERROR: "Client Error",
      },
      text: {
        CONNECTING: "Connecting to Guacamole...",
        DISCONNECTED: "You have been disconnected.",
        UNSTABLE:
          "The network connection to the Guacamole server appears unstable.",
        WAITING: "Connected to Guacamole. Waiting for response...",
      },
    };
  },
  computed: {
    canReconnect() {
      return ["DISCONNECTED", "CLIENT_ERROR"].includes(this.status);
    },
  },
  created() {
    console.log("머라는거여???????????????????");
  },
  methods: {
    show(state, message) {
      console.log(state, message);
      if (state === states.CONNECTED) {
        this.status = null;
      } else {
        this.status = state;
      }
      this.message = message;
    },
  },
};
</script>
<style scoped>
.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
  padding: 1rem;
}
.rct {
  text-decoration: underline;
  cursor: pointer;
}
</style>
