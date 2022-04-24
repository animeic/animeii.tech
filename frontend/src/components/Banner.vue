<template>
    <div id="banner-container">
        <div class="item-1"><a href="" @click="clickhandle">一月番</a></div>
        <div class="item-4"><a href="">四月番</a></div>
        <div class="item-7"><a href="">七月番</a></div>
        <div class="item-10"><a href="">十月番</a></div>
    </div>
</template>

<style lang="scss" scoped>
#banner-container{
   // grid
   display: grid;
   height: 120px;
   grid-template-columns: repeat(4,1fr);
   justify-items: center; // 水平对齐
   align-items: center;
   // debug
   border-bottom: 1px solid gray;

}
</style>

<script setup>

const clickhandle = ()=>{
    alert("aaa")
}

</script>