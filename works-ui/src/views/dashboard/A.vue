<template>
  <div class="dashboard-main-grid">
    <div class="dashboard-left-grid">
      <a-card :loading="loading" class="dashboard-card">
        <h3>워크스페이스 수</h3>
        <a-progress
            type="dashboard"
            :percent="70"
            :size="large" />
        <p>20 / 30</p>
      </a-card>
      <a-card :loading="loading" class="dashboard-card">
        <h3>워크스페이스 수</h3>
        <a-progress
            type="dashboard"
            :percent="70"
            :size="large" />
        <p>20 / 30</p>
      </a-card>
      <a-card :loading="loading" class="dashboard-card">
        <h3>워크스페이스 수</h3>
        <a-progress
            type="dashboard"
            :percent="70"
            :size="large" />
        <p>20 / 30</p>
      </a-card>
      <a-card :loading="loading" class="dashboard-card">
        <h3>워크스페이스 수</h3>
        <a-progress
            type="dashboard"
            :percent="70"
            :size="large" />
        <p>20 / 30</p>
      </a-card>
      <a-card :loading="loading" class="dashboard-card">
        <h3>워크스페이스 수</h3>
        <a-progress
            type="dashboard"
            :percent="70"
            :size="large" />
        <p>20 / 30</p>
      </a-card>
      <a-card :loading="loading" class="dashboard-card">
        <h3>워크스페이스 수</h3>
        <a-progress
            type="dashboard"
            :percent="70"
            :size="large" />
        <p>20 / 30</p>
      </a-card>
    </div>
    <div class="dashboard-right-grid">
      <a-card title="공지사항 제목" :bordered="true" class="dashboard-right-card">
        <p style="font-weight: bold">작성자 : ABLECLOUD</p>
        <p class="dashboard-right-card-p">공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. </p>
      </a-card>
      <a-card title="공지사항 제목" :bordered="true" class="dashboard-right-card">
        <p style="font-weight: bold">작성자 : ABLECLOUD</p>
        <p class="dashboard-right-card-p">공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. </p>
      </a-card>
      <a-card title="공지사항 제목" :bordered="true" class="dashboard-right-card">
        <p style="font-weight: bold">작성자 : ABLECLOUD</p>
        <p class="dashboard-right-card-p">공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다. </p>
      </a-card>
    </div>
  </div>
</template>

<script>
export default {
  name: "A"
}
</script>

<style scoped>
.dashboard-main-grid{
  display: grid;
  grid-template-columns: 3fr 1fr;
  grid-template-rows: 1fr;
  height: 100%;
  text-align: center;
  padding: 5px;
}
.dashboard-left-grid{
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  grid-template-rows: 291px 291px;
  grid-gap: 14px;
  padding: 14px;
  text-align: center;
}

.dashboard-right-grid{
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr 1fr 1fr;
  grid-gap: 14px;
  padding: 14px;
}

.dashboard-card{
  max-height: 291px;
}
.dashboard-right-card{
  width: 100%;
  text-align: left;
  margin-bottom: 10px;
  box-shadow: 5px 5px;
}
.dashboard-right-card-p{
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  line-height: 16px;
  max-height: 50px;
  -webkit-line-clamp: 3; /* 표시하고자 하는 라인 수 */
  -webkit-box-orient: vertical;
}
</style>
