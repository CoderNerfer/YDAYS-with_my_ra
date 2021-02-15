$(document).ready(function(){
    $(window).scroll(function(){
        if(this.scrollY > 20){
            $(".navbar").addClass("sticky");
            $(".goTop").fadeIn();
        }
        else{
            $(".navbar").removeClass("sticky");
            $(".goTop").fadeOut();
        }
        
    });

    $(".goTop").click(function(){scroll(0,0)});

    $('.menu-toggler').click(function(){
        $(this).toggleClass("active");
        $(".navbar-menu").toggleClass("active");
    });

    document.getElementById('copybutton').addEventListener('click', function(){
        document.querySelector('.bg-modal').style.display = 'flex';
    });
    document.getElementById('copybutton').addEventListener('click', function(){
        navigator.clipboard.writeText('withmyra2@gmail.com');
        console.log('done');
    });

    document.querySelector('.close').addEventListener('click',function(){
        document.querySelector('.bg-modal').style.display = 'none';
    });

});
