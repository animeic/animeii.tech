<template>
<div id="footer-container">
    <div class="logo"></div>
    <div class="search"></div>
    <div class="item1"></div>
    <div class="item2"></div>
    <div class="item3"></div>
    <div class="login"></div>
    <div class="register"></div>
</div>
</template>

<style lang="scss" scoped>
#footer-container{
     // main
    height: 56px;
    background-color: #e6bce3;
    // grid
    display: grid;
    grid-template-columns: repeat(9,1fr);
    grid-template-areas: 'logo search search search item1 item2 item3 login register';
    place-items: center;
    // items
    .logo{
        grid-area: logo;
        font-size: large;
    }
    .search{
        grid-area: search;
    }
    .item1{
        grid-area: item1;
    }
    .item2{
        grid-area: item2;
    }
    .item3{
        grid-area: item3;
    }
    .login{
        grid-area: login;
    }
    .register{
        grid-area: register;
    }

}
</style>