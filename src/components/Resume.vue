<template>
  <div class="resume-container">
    <!-- 头部信息 -->
    <header class="resume-header">
      <div class="header-content">
        <div class="avatar-container">
          <img src="https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124225313.png" alt="头像" class="avatar" />
        </div>
        <h1 class="name">个人在线简历</h1>
        <p class="target">求职意向：后端开发工程师 | 期望薪资：税前8K</p>
        <div class="contact-info">
          <div class="contact-item">
            <span class="label">邮箱：</span>
            <a href="mailto:18319714791@163.com" class="value">18319714791@163.com</a>
          </div>
          <div class="contact-item">
            <span class="label">微信：</span>
            <span class="value">Rolin3456</span>
          </div>
        </div>
      </div>
    </header>

    <!-- 主要内容 -->
    <main class="resume-main">
      <!-- 工作经历 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">💼</span>
          工作经历
        </h2>
        
        <div class="experience-item">
          <div class="experience-header">
            <h3 class="company">广州小鸡快跑网络科技有限公司</h3>
            <span class="position">后端开发工程师</span>
            <span class="period">2025.7 - 2025.11</span>
          </div>
          <p class="project-desc">
            <strong>项目描述：</strong>基于 Go-Zero 微服务框架的游戏平台后端系统，采用 Monorepo 单体仓库架构，支持游戏管理、用户认证、资源管理、Steam 集成等功能，使用 PostgreSQL + Redis 作为数据存储，支持多语言国际化。
          </p>
          <ul class="work-content">
            <li>基于 PostgreSQL + Jieba 中文分词实现全文搜索，使用 GIN 索引优化查询性能至 100ms 内，设计智能权重排序算法，搜索结果相关性提升 40%。</li>
            <li>设计并实现 Game、OAuth、Resource、CDC 四大微服务，制定三层架构规范，通过泛型封装 BaseDAO 实现 CRUD 通用方法与工具类，减少 80% 重复代码，提升代码复用性。</li>
            <li>实现 CDC 监听数据库变更自动更新缓存，基于 Redis Stream + 多 worker 并行消费，支持 1000+ TPS。</li>
            <li>封装 errorx 统一错误处理机制，所有异常携带详细上下文信息。实现 OAuth2.0 + JWT 认证体系，支持多平台第三方登录，集成支付宝/微信支付。</li>
          </ul>
          <div class="departure-note">
            <p><strong>说明：</strong>在职期间表现优异，已成功转正并获得公司认可。因公司进行架构调整，业务方向发生变化，遭遇裁员而离职，非个人能力原因。</p>
          </div>
        </div>
      </section>

      <!-- 实习经历 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">🎓</span>
          实习经历
        </h2>
        
        <div class="experience-item">
          <div class="experience-header">
            <h3 class="company">珠海一讯牵通信科技有限公司</h3>
            <span class="position">后端开发(实习)</span>
            <span class="period">2024.9 - 2025.3</span>
          </div>
          <p class="project-name"><strong>项目名称：</strong>智慧规建</p>
          <p class="project-desc">
            <strong>项目描述：</strong>基于分布式Spring架构开发的便民小程序，总分为运营后台和公众平台两大模块，具有个人中心、后台管理、预约互动、积分互动、抽奖活动、城市主题、城市活动、城市咨询、城市一张图等功能。
          </p>
          <ul class="work-content">
            <li>参与敲定预约模块、城市一张图的整体架构与交互逻辑，排除可能会出现的潜在问题。</li>
            <li>开发同步咨询接口，使用线程池管理多个请求线程，解决请求过多时占用大量资源问题。开发针对失败请求的重试功能，减少全量请求的同步次数。开发错误日志记录功能，保证任务失败时操作人能看到失败原因。</li>
            <li>编写多个功能的详细设计文档、PDM、修复Bug，参与项目详细设计评审和功能提测，持续推进项目完成。</li>
            <li>编写生成数据脚本与基础工具接口，减少测试人员的工作，减少重复代码并提高项目开发体验。</li>
          </ul>
        </div>

        <div class="experience-item">
          <div class="experience-header">
            <h3 class="company">星园科技（广州）有限公司</h3>
            <span class="position">Java开发(实习)</span>
            <span class="period">2023.7 - 2023.11</span>
          </div>
          <p class="project-name"><strong>项目名称：</strong>微信小程序搜索Hi校园</p>
          <p class="project-desc">
            <strong>项目描述：</strong>该项目已上线，主要用于校园学生使用，相关业务包括课程查询、成绩查询、个人信息查询等功能，使用人数主要为各个学校的在校生2万多人。
          </p>
          <ul class="work-content">
            <li>在爬虫模块中，使用 HtmlUnit、Selenium、HttpClient 等框架根据用户账号密码登入学生个人信息中心通过元素定位查找目标信息并存入 MySQL 数据库中，方便信息的重复使用。</li>
            <li>优化已存在的爬虫代码，将原使用的重量级框架 Selenium 替换为轻量级框架 HtmlUnit，降低远程部署时需要的内存并提升爬虫运行效率。</li>
            <li>通过提取爬虫重复代码来制定通用的爬虫接口并参与开发文档的编写，大幅降低了新加入公司的程序员的进行爬虫所需的业务了解时间。</li>
            <li>开发武汉体、武汉科、武汉师等学校的爬虫并将其部署到远程服务器中，实现自动爬虫。</li>
          </ul>
        </div>
      </section>

      <!-- 在校经历 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">📚</span>
          在校经历
        </h2>
        <div class="education-content">
          <p>大一期间专注于Java技术栈的学习与实践，通过刷力扣算法题提升编程能力和算法思维。大二加入省重点实验室，积极参与多个科研项目的研发工作，其中包括农产品大数据平台和柑橘智慧云平台等创新项目。</p>
          <p>在校期间获得大创校级资助，并在蓝桥杯省赛中荣获二等奖，英语通过四级考试。通过在校学习和实习历练，培养了独立开发项目的能力、良好的团队沟通协作能力，以及较强的抗压能力和快速适应新环境的能力。</p>
        </div>
      </section>

      <!-- 个人荣誉 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">🏆</span>
          个人荣誉
        </h2>
        <ul class="honor-list">
          <li>2024年10月申报获得"基于物联网与人工智能的柑橘智慧云平台V1.0"的计算机软件著作权登记证书。</li>
          <li>2024年3月申报"基于物联网和人工智能的柑橘智慧云服务平台"大创项目获得校级资助。</li>
          <li>2023年11月获得TiDB Associate职位证书。</li>
        </ul>
      </section>

      <!-- 专业技能 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">🛠️</span>
          专业技能
        </h2>
        <ul class="skill-list">
          <li><strong>Java技术栈：</strong>熟练掌握Java语法、集合、多线程等基础框架原理（synchronized、ThreadLocal、HashMap、线程池）</li>
          <li><strong>分布式与微服务：</strong>熟悉分布式基本理论知识和微服务相关框架的使用开发（CAP、BASE、Redis分布式锁、Sentinel、RabbitMQ）</li>
          <li><strong>JVM与并发：</strong>熟悉JVM和JUC的基本知识（GC算法、常见的垃圾回收器、类加载过程、Volatile、CAS、AQS）</li>
          <li><strong>Redis：</strong>熟练非关系型数据库Redis（数据结构、高可用、性能调优）</li>
          <li><strong>MySQL：</strong>熟练使用关系型数据库MySQL（索引、SQL语句调优、事务、锁、新版本特性）</li>
          <li><strong>Go语言：</strong>熟悉Go语言基础语法和特性，了解Go-Zero微服务框架，具备Go语言后端开发经验</li>
          <li><strong>Go并发编程：</strong>熟悉Goroutine、Channel、Context等Go并发编程机制，了解Go的内存模型和调度器原理</li>
          <li><strong>Go生态：</strong>熟悉Gin、GORM等常用框架，了解Go的依赖管理和模块化开发</li>
          <li><strong>Linux：</strong>熟悉Linux常用命令，会使用Linux部署项目，掌握基本的数据结构和算法</li>
        </ul>
      </section>

      <!-- 联系方式 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">📞</span>
          联系我
        </h2>
        <div class="contact-section">
          <div class="contact-text">
            <p>个人微信：<strong>Rolin3456</strong></p>
            <p class="contact-note">如有内推机会，请通过该微信联系我</p>
            <p>邮箱：<a href="mailto:18319714791@163.com" class="email-link">18319714791@163.com</a></p>
          </div>
          <div class="wechat-qr">
            <img 
              src="https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/893a475465e3f4f1a0bcfb015caeea74.jpg" 
              alt="微信二维码" 
              class="qr-image"
              @click="showQrModal = true"
            />
            <p class="qr-tip">扫码添加微信（点击放大）</p>
          </div>
        </div>
      </section>

      <!-- 证明材料 -->
      <section class="resume-section">
        <h2 class="section-title">
          <span class="title-icon">📄</span>
          证明材料
        </h2>
        <div class="certificate-grid">
          <div class="certificate-item" @click="openImageModal('https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124230232.png', '学信网截图')">
            <div class="certificate-image-wrapper">
              <img 
                src="https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124230232.png" 
                alt="学信网截图" 
                class="certificate-image"
              />
              <div class="certificate-overlay">
                <span class="certificate-icon">🔍</span>
                <span class="certificate-text">点击查看大图</span>
              </div>
            </div>
            <p class="certificate-name">学信网截图</p>
          </div>
          <div class="certificate-item" @click="openImageModal('https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124230644.png', 'TiDB Associate证书')">
            <div class="certificate-image-wrapper">
              <img 
                src="https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124230644.png" 
                alt="TiDB Associate证书" 
                class="certificate-image"
              />
              <div class="certificate-overlay">
                <span class="certificate-icon">🔍</span>
                <span class="certificate-text">点击查看大图</span>
              </div>
            </div>
            <p class="certificate-name">TiDB Associate证书</p>
          </div>
          <div class="certificate-item" @click="openImageModal('https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124230716.png', '蓝桥杯省赛第二名证书')">
            <div class="certificate-image-wrapper">
              <img 
                src="https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/20251124230716.png" 
                alt="蓝桥杯省赛第二名证书" 
                class="certificate-image"
              />
              <div class="certificate-overlay">
                <span class="certificate-icon">🔍</span>
                <span class="certificate-text">点击查看大图</span>
              </div>
            </div>
            <p class="certificate-name">蓝桥杯省赛第二名证书</p>
          </div>
        </div>
      </section>
    </main>

    <!-- 页脚备案信息 -->
    <footer class="resume-footer">
      <a href="https://beian.miit.gov.cn/#/Integrated/index" target="_blank" class="beian-link">
        粤ICP备2025451803号-1
      </a>
    </footer>

    <!-- 二维码放大模态框 -->
    <div v-if="showQrModal" class="modal-overlay" @click="showQrModal = false">
      <div class="modal-content" @click.stop>
        <button class="modal-close" @click="showQrModal = false">&times;</button>
        <img 
          src="https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/893a475465e3f4f1a0bcfb015caeea74.jpg" 
          alt="微信二维码" 
          class="modal-qr-image"
        />
        <p class="modal-qr-tip">微信二维码</p>
      </div>
    </div>

    <!-- 图片放大模态框 -->
    <div v-if="showImageModal" class="modal-overlay" @click="closeImageModal">
      <div class="modal-content" @click.stop>
        <button class="modal-close" @click="closeImageModal">&times;</button>
        <img 
          :src="modalImageSrc" 
          :alt="modalImageAlt" 
          class="modal-image"
        />
        <p class="modal-qr-tip">{{ modalImageAlt }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'Resume',
  setup() {
    const showQrModal = ref(false)
    const showImageModal = ref(false)
    const modalImageSrc = ref('')
    const modalImageAlt = ref('')

    const openImageModal = (src, alt) => {
      modalImageSrc.value = src
      modalImageAlt.value = alt
      showImageModal.value = true
    }

    const closeImageModal = () => {
      showImageModal.value = false
      modalImageSrc.value = ''
      modalImageAlt.value = ''
    }

    return {
      showQrModal,
      showImageModal,
      modalImageSrc,
      modalImageAlt,
      openImageModal,
      closeImageModal
    }
  }
}
</script>

<style scoped>
.resume-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 40px 20px;
  background: #ffffff;
  min-height: 100vh;
  box-shadow: 0 0 30px rgba(0, 0, 0, 0.1);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
}

.resume-header {
  text-align: center;
  padding: 40px 0;
  border-bottom: 3px solid #3498db;
  margin-bottom: 40px;
}

.avatar-container {
  margin-bottom: 20px;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  border: 4px solid #3498db;
  object-fit: cover;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.header-content .name {
  font-size: 36px;
  font-weight: 700;
  color: #2c3e50;
  margin: 0 0 15px 0;
}

.header-content .target {
  font-size: 18px;
  color: #555;
  margin: 0 0 20px 0;
}

.contact-info {
  display: flex;
  justify-content: center;
  gap: 30px;
  flex-wrap: wrap;
}

.contact-item {
  font-size: 16px;
}

.contact-item .label {
  color: #666;
}

.contact-item .value {
  color: #3498db;
  text-decoration: none;
}

.contact-item .value:hover {
  text-decoration: underline;
}

.resume-section {
  margin-bottom: 50px;
}

.section-title {
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 25px;
  padding-bottom: 10px;
  border-bottom: 2px solid #ecf0f1;
  display: flex;
  align-items: center;
  gap: 10px;
}

.title-icon {
  font-size: 28px;
}

.experience-item {
  margin-bottom: 35px;
  padding: 25px;
  background: #f8f9fa;
  border-radius: 8px;
  border-left: 4px solid #3498db;
}

.experience-header {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-bottom: 15px;
  flex-wrap: wrap;
}

.company {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.position {
  padding: 5px 12px;
  background: #3498db;
  color: white;
  border-radius: 4px;
  font-size: 14px;
}

.period {
  color: #fff;
  font-size: 16px;
  font-weight: 600;
  margin-left: auto;
  padding: 8px 16px;
  background: linear-gradient(135deg, #3498db 0%, #2980b9 100%);
  border-radius: 20px;
  box-shadow: 0 2px 8px rgba(52, 152, 219, 0.3);
  white-space: nowrap;
}

.project-name {
  margin: 10px 0;
  color: #555;
}

.project-desc {
  margin: 15px 0;
  line-height: 1.8;
  color: #555;
}

.work-content {
  margin: 15px 0 0 0;
  padding-left: 20px;
}

.work-content li {
  margin-bottom: 12px;
  line-height: 1.8;
  color: #555;
}

.departure-note {
  margin-top: 20px;
  padding: 15px 20px;
  background: #f0f7ff;
  border-left: 4px solid #3498db;
  border-radius: 6px;
  line-height: 1.8;
}

.departure-note p {
  margin: 0;
  color: #555;
  font-size: 14px;
}

.departure-note strong {
  color: #2c3e50;
}

.education-content {
  line-height: 1.8;
  color: #555;
}

.education-content p {
  margin-bottom: 15px;
}

.honor-list,
.skill-list {
  list-style: none;
  padding: 0;
}

.honor-list li,
.skill-list li {
  margin-bottom: 15px;
  padding: 15px;
  background: #f8f9fa;
  border-radius: 6px;
  line-height: 1.8;
  color: #555;
  border-left: 3px solid #3498db;
}

.skill-list li strong {
  color: #2c3e50;
  display: inline-block;
  min-width: 120px;
}

.contact-section {
  display: flex;
  gap: 40px;
  align-items: flex-start;
  flex-wrap: wrap;
}

.contact-text {
  flex: 1;
  min-width: 250px;
}

.contact-text p {
  margin-bottom: 12px;
  line-height: 1.8;
  color: #555;
}

.contact-note {
  color: #e74c3c;
  font-size: 14px;
  font-style: italic;
}

.email-link {
  color: #3498db;
  text-decoration: none;
}

.email-link:hover {
  text-decoration: underline;
}

.wechat-qr {
  text-align: center;
  flex-shrink: 0;
}

.wechat-qr img {
  width: 180px;
  height: 180px;
  border: 2px solid #ecf0f1;
  border-radius: 8px;
  display: block;
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.qr-image:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(52, 152, 219, 0.3);
}

.qr-tip {
  margin-top: 10px;
  font-size: 14px;
  color: #7f8c8d;
}

.resume-footer {
  text-align: center;
  padding: 30px 0;
  margin-top: 50px;
  border-top: 1px solid #ecf0f1;
}

.beian-link {
  color: #95a5a6;
  text-decoration: none;
  font-size: 14px;
}

.beian-link:hover {
  color: #3498db;
  text-decoration: underline;
}

/* 模态框样式 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

.modal-content {
  position: relative;
  background: white;
  padding: 30px;
  border-radius: 12px;
  text-align: center;
  max-width: 90%;
  max-height: 90%;
  animation: scaleIn 0.3s ease;
}

.modal-close {
  position: absolute;
  top: 10px;
  right: 15px;
  background: none;
  border: none;
  font-size: 32px;
  color: #7f8c8d;
  cursor: pointer;
  line-height: 1;
  padding: 0;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.3s ease;
}

.modal-close:hover {
  color: #e74c3c;
}

.modal-qr-image {
  max-width: 400px;
  max-height: 400px;
  width: auto;
  height: auto;
  border: 2px solid #ecf0f1;
  border-radius: 8px;
  display: block;
  margin: 0 auto;
}

.modal-qr-tip {
  margin-top: 15px;
  font-size: 16px;
  color: #555;
}

/* 证明材料样式 */
.certificate-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 30px;
  margin-top: 20px;
}

.certificate-item {
  cursor: pointer;
  transition: transform 0.3s ease;
}

.certificate-item:hover {
  transform: translateY(-5px);
}

.certificate-image-wrapper {
  position: relative;
  width: 100%;
  padding-top: 75%;
  background: #f8f9fa;
  border-radius: 8px;
  overflow: hidden;
  border: 2px solid #ecf0f1;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
}

.certificate-item:hover .certificate-image-wrapper {
  border-color: #3498db;
  box-shadow: 0 4px 12px rgba(52, 152, 219, 0.2);
}

.certificate-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.certificate-item:hover .certificate-image {
  transform: scale(1.05);
}

.certificate-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(52, 152, 219, 0.85);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
  color: white;
}

.certificate-item:hover .certificate-overlay {
  opacity: 1;
}

.certificate-icon {
  font-size: 32px;
  margin-bottom: 8px;
}

.certificate-text {
  font-size: 14px;
  font-weight: 500;
}

.certificate-name {
  margin-top: 12px;
  text-align: center;
  font-size: 16px;
  font-weight: 500;
  color: #2c3e50;
}

.modal-image {
  max-width: 90vw;
  max-height: 85vh;
  width: auto;
  height: auto;
  border: 2px solid #ecf0f1;
  border-radius: 8px;
  display: block;
  margin: 0 auto;
  object-fit: contain;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes scaleIn {
  from {
    transform: scale(0.8);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}

@media (max-width: 768px) {
  .resume-container {
    padding: 20px 15px;
    margin: 0;
    box-shadow: none;
  }

  .header-content .name {
    font-size: 28px;
  }

  .header-content .target {
    font-size: 16px;
  }

  .contact-info {
    flex-direction: column;
    gap: 10px;
  }

  .experience-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .period {
    margin-left: 0;
  }

  .contact-section {
    flex-direction: column;
  }

  .wechat-qr {
    align-self: center;
  }

  .section-title {
    font-size: 20px;
  }

  .experience-item {
    padding: 20px 15px;
  }

  .avatar {
    width: 100px;
    height: 100px;
  }

  .modal-content {
    padding: 20px;
  }

  .modal-qr-image {
    max-width: 300px;
    max-height: 300px;
  }

  .certificate-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .period {
    font-size: 14px;
    padding: 6px 12px;
    margin-left: 0;
    margin-top: 10px;
  }
}
</style>
