<template>
  <div class="container">
    <!-- 关于我 -->
    <AboutMe />
    
    <!-- 特别鸣谢 -->
    <Acknowledgments />
    
    <!-- 目前更新计划 -->
    <div class="update-plan">
      <h3>更新计划说明</h3>
      <p>工作日每天更新一期无解说对局视频，若有推广合作会更新解说局。</p>
      <p>当前状态将持续至国庆假期，期间暂不更新常规解说局及其他内容。</p>
      <p>国庆后将根据工作转正情况调整更新计划，若顺利转正将投入更多时间运营账号。</p>
    </div>

    <!-- 对局投稿 (已暂时隐藏) -->
    <!-- <SubmissionForm @submission-submitted="handleSubmission" /> -->
    
    <!-- 交流群 -->
    <Community />
    
    <!-- 支持主包 -->
    <Support />
    
    <!-- 粉丝福利 -->
    <FanBenefits />
    
    <!-- 必看视频 -->
    <MustWatch />
    
    <!-- 页脚 -->
    <footer class="footer">
      <a href="http://beian.miit.gov.cn/" target="_blank" rel="noopener noreferrer">
        粤ICP备2025451803号-1
      </a>
    </footer>
  </div>
</template>

<script>
import AboutMe from './AboutMe.vue'
import Acknowledgments from './Acknowledgments.vue'
import SubmissionForm from './SubmissionForm.vue'
import Community from './Community.vue'
import Support from './Support.vue'
import FanBenefits from './FanBenefits.vue'
import MustWatch from './MustWatch.vue'

export default {
  name: 'UserFront',
  components: {
    AboutMe,
    Acknowledgments,
    SubmissionForm,
    Community,
    Support,
    FanBenefits,
    MustWatch
  },
  methods: {
    handleSubmission(submission) {
      // 将投稿数据发送到全局事件总线，供后台使用
      this.$emit('submission-submitted', submission)
      
      // 这里可以添加本地存储逻辑，如果需要持久化的话
      // localStorage.setItem('submissions', JSON.stringify([...existingSubmissions, submission]))
    }
  }
}
</script>