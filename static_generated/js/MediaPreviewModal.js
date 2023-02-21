var a=Object.defineProperty;var c=(s,i,e)=>i in s?a(s,i,{enumerable:!0,configurable:!0,writable:!0,value:e}):s[i]=e;var t=(s,i,e)=>(c(s,typeof i!="symbol"?i+"":i,e),e);const n=`
<link rel="stylesheet" href="/static/css/fileshare.css">
<dialog class="media-preview-modal">
    <article class="media-preview-modal-article">
        <a href="#" aria-label="Close" id="form-modal-close" class="close"></a>
        <div id="media-preview-slot">
          <slot name="media"></slot>
        </div>
    </article>
</dialog>
`;class o extends HTMLDialogElement{constructor(){super();t(this,"template");t(this,"content");t(this,"closeButton");t(this,"dialog");t(this,"article");t(this,"visible",!1);t(this,"mediaRoot");t(this,"persist",!1);t(this,"animateInOut",!1);t(this,"closeButtonListener",()=>{this.closeModal()});t(this,"clickOffArticleListener",e=>{if(this.persist||!this.visible)return;const l=this.article.getBoundingClientRect();e.clientX>=l.left&&e.clientX<=l.right&&e.clientY>=l.top&&e.clientY<=l.bottom||this.closeModal()});this.template=document.createElement("template"),this.template.innerHTML=n,this.content=this.template.content;const e=this.attachShadow({mode:"open"});e.appendChild(this.content.cloneNode(!0)),this.dialog=e.querySelector("dialog"),this.article=e.querySelector("article"),this.closeButton=e.querySelector("#form-modal-close"),this.mediaRoot=e.querySelector("#media-preview-slot")}closeModal(){if(!this.animateInOut){this.dialog.close(),this.visible=!1,this.mediaRoot.replaceChildren();return}this.article.style.animation="fadeOutUp 0.3s forwards",setTimeout(()=>{this.dialog.close(),this.visible=!1,this.mediaRoot.replaceChildren()},300)}showModal(){if(setTimeout(()=>{this.visible=!0},100),!this.animateInOut){this.dialog.showModal();return}this.article.style.animation="fadeInDown 0.3s forwards",this.dialog.showModal()}connectedCallback(){this.persist=this.hasAttribute("persist"),this.animateInOut=this.hasAttribute("animate"),this.closeButton.addEventListener("click",this.closeButtonListener),document.addEventListener("keydown",e=>{e.key==="Escape"&&this.visible&&this.closeModal()}),document.addEventListener("click",this.clickOffArticleListener)}disconnectedCallback(){this.closeButton.removeEventListener("click",this.closeButtonListener),document.removeEventListener("click",this.clickOffArticleListener)}static register(){customElements.define("media-preview-modal",o,{extends:"dialog"})}}export{o as M};