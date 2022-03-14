import {
    Star
} from '../mymodule/stars.js'
import {
    Updown
} from '../mymodule/updown.js'
let moviecommendid = sessionStorage.getItem('moviecommendid');
let discussid = sessionStorage.getItem('discussid')
let mytext = document.querySelector('.mytext');
let responsearea = document.querySelector('.responsearea')
let discussarea = document.querySelector('.discussarea');
let mytoken = sessionStorage.getItem('token')

function getmoviecommend(url) {
    return new Promise(resolve => {
        fetch(url, {
                method: 'GET'
            }).then(res => res.json())
            .then(res => {
                console.log(res.data.id);
                mytext.innerHTML =
                    `<h1>${res.data.context.split(/[，。？&]/)[0]}</h1>
        <div class="info">
            <span class="peoplename">${res.data.name}</span>
            <span>评论</span>
            <span class="moviename">${res.data.movie_name}</span>
            <div id="stars"></div>
            <span class="date">${res.data.post_time}</span>
        </div>
        <div class="commendcontext">${res.data.context.split('\n').map(item=>{
            return `<p>${item}</p>`
        }).join('')}</div>`
                let mystar = new Star('stars', res.data.star_num, 0.5);
                mystar.create();
                return res;
            }).then(res => resolve(res))
    })
}

function toaddres(url, myformdata, mytoken) {
    fetch(url, {
            method: 'POST',
            body: myformdata,
            headers: {
                token: mytoken
            }
        }).then(res => res.json())
        .then(res => {
            alert(res.info)
            window.location.reload();
        })
}
window.onload = async function () {
    let respond = document.querySelector('textarea');
    let toadd = document.querySelector('button');
    if (sessionStorage.getItem('isdiscuss') == 'false') {
        responsearea.style.display = 'block';
        discussarea.style.display = 'none';
        let res1 = await getmoviecommend('http://42.192.155.29:8080/filmcomment/' + moviecommendid)
        res1.data.FilmCommentReplys.map(item => {
            responsearea.innerHTML +=
                `<li>
    <img src="./tiger1.png" alt="">
    <div class="mycontext">
        <div class="meta-header">
            <a href="#">${item.Name}</a>
            <span>${item.CommentTime}</span>
        </div>
        <div class="mybody">${item.Context}</div>
    </div>
</li>`
        })
        toadd.addEventListener('click', () => {
            let formdata = new FormData();
            formdata.append('context', respond.value)
            if (mytoken) {
                toaddres('http://42.192.155.29:8080/filmcomment_reply/' + moviecommendid, formdata, mytoken)
            } else {
                alert("请先登录豆瓣")
            }
        })
    } else {
        responsearea.style.display = 'none';
        discussarea.style.display = 'block';
        let res2 = await getmoviecommend('http://42.192.155.29:8080/topic/' + discussid)
        res2.data.Comments.map(item => {
            console.log(item);
            discussarea.innerHTML +=
                `<li>
    <img src="./tiger1.png" alt="">
    <div class="mycontext">
        <div class="meta-header" id="head2">
            <a href="#">${item.Name}</a>
            <span>${item.CommentTime}</span>
        </div>
        <div class="mybody">${item.Context}</div>
        <div class="oplinks">
        <span>来自 豆瓣APP</span>
        <span class="myups">
        赞=>
        <span class="changeup">${item.Likes}</span>
        </span>
        </div>
    </div>
</li>`
        })
        let myups=document.querySelectorAll('.myups');
        let changeup=document.querySelectorAll('.changeup')
        res2.data.Comments.map((item,index)=>{
            let up= new Updown(myups[index],changeup[index],item.Likes,'http://42.192.155.29:8080/comment/likes/'+item.Id);
            up.add();
        })
        toadd.addEventListener('click', () => {
            let formdata = new FormData();
            formdata.append('context', respond.value)
            if (mytoken) {
                toaddres('http://42.192.155.29:8080/comment/' + discussid, formdata, mytoken)
            } else {
                alert("请先登录豆瓣")
            }
        })
    }
}