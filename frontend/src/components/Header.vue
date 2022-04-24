<template>
<div id="header-container">
    <div class="logo"><el-link href="/" >animeic</el-link></div>
    <Search/>
    <div class="article"><el-link href="/article">文章</el-link></div>
    <div class="register"><el-link href="/register">注册</el-link></div>
</div>
</template>

<script setup>
import Search from '../components/Search.vue'
</script>

<style lang="scss" scoped>
#header-container{
     // main
    height: 60px;
    border-bottom: 1px solid #dcdfe6;
    // grid
    display: grid;
    grid-template-columns: repeat(9,1fr);
    place-items: center;

    .logo{
        grid-area: 1/2/2/3;   
        a{
            font-size: 20px;
            color: hotpink;
            text-decoration: none;
        }
    }
    // 子组件的容器 这些属性必须。子组件中在渲染前并不制定有这些属性
    #search-container{
        // grid
        grid-area: 1/4/2/7;
        justify-self: stretch;

    }
    .article{
        grid-area: 1 / 7 / 2 / 9
       
    }
    .register{
        grid-area: 1 / 9 / 2 / 10;
       
    }
     a{
        color: hotpink;
        text-decoration: none;
    }
    

}
</style>