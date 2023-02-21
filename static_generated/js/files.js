var y=Object.defineProperty;var L=(t,s,e)=>s in t?y(t,s,{enumerable:!0,configurable:!0,writable:!0,value:e}):t[s]=e;var o=(t,s,e)=>(L(t,typeof s!="symbol"?s+"":s,e),e);import{M as w}from"./MediaPreviewModal.js";const B=`
	<link rel="stylesheet" href="/static/css/fileshare.css">
	<dialog id="confirm-modal">
		<article id="confirm-article">
			<a href="#" aria-label="Close" id="confirm-modal-close" class="close"></a>
			<h3>
				<slot id="confirm-modal-title" name="title">Title</slot>
			</h3>
			<p>
				<slot id="confirm-modal-message" name="message">Are you sure you want to do this?</slot>
			</p>
			<footer id="confirm-modal-action-footer">
				<button id="confirm-modal-cancel-button">Cancel</button>
				<button id="confirm-modal-confirm-button">Confirm</button>
			</footer>
		</article>
	</dialog>`;class g extends HTMLDialogElement{constructor(){super();o(this,"template");o(this,"content");o(this,"closeButton");o(this,"confirmButton");o(this,"cancelButton");o(this,"dialog");o(this,"article");o(this,"animateInOut",!1);o(this,"eventData");o(this,"visible",!1);o(this,"persistEventHandler",null);o(this,"titleSlotElement");o(this,"messageSlotElement");o(this,"closeModal",()=>{if(!this.animateInOut){this.dialog.close(),this.visible=!1;return}this.article.style.animation="fadeOutUp 0.3s forwards",setTimeout(()=>{this.dialog.close(),this.visible=!1},300)});o(this,"showModal",()=>{if(setTimeout(()=>{this.visible=!0},100),!this.animateInOut){this.dialog.showModal();return}this.article.style.animation="fadeInDown 0.3s forwards",this.dialog.showModal()});o(this,"closeButtonListener",()=>{this.dispatchEvent(new CustomEvent("cancel",{detail:this.eventData})),this.closeModal()});o(this,"cancelButtonListener",()=>{this.dispatchEvent(new CustomEvent("cancel",{detail:this.eventData})),this.closeModal()});o(this,"confirmButtonListener",()=>{this.dispatchEvent(new CustomEvent("confirm",{detail:this.eventData})),this.closeModal()});o(this,"noPersistHandler",e=>{this.visible&&(this.article.contains(e.target)||this.closeModal())});this.template=document.createElement("template"),this.template.innerHTML=B,this.content=this.template.content;const e=this.attachShadow({mode:"open"});e.appendChild(this.content.cloneNode(!0)),this.closeButton=e.querySelector("#confirm-modal-close"),this.confirmButton=e.querySelector("#confirm-modal-confirm-button"),this.cancelButton=e.querySelector("#confirm-modal-cancel-button"),this.dialog=e.querySelector("#confirm-modal"),this.article=e.querySelector("#confirm-article"),this.titleSlotElement=e.querySelector("#confirm-modal-title"),this.messageSlotElement=e.querySelector("#confirm-modal-message")}static get observedAttributes(){return["animate","persist","confirm-button-class","cancel-button-class","event-data","title","message"]}connectedCallback(){var i,n;this.closeButton.addEventListener("click",this.closeButtonListener),this.confirmButton.addEventListener("click",this.confirmButtonListener),this.cancelButton.addEventListener("click",this.cancelButtonListener),((i=this.getAttribute("persist"))==null?void 0:i.toLowerCase())!=="true"&&document.documentElement.addEventListener("click",this.noPersistHandler),this.animateInOut=((n=this.getAttribute("animate"))==null?void 0:n.toLowerCase())==="true",this.eventData=this.getAttribute("event-data")??null,this.confirmButton.className=this.getAttribute("confirm-button-class")??"primary",this.cancelButton.className=this.getAttribute("cancel-button-class")??"secondary"}disconnectedCallback(){this.closeButton.removeEventListener("click",this.closeButtonListener),this.confirmButton.removeEventListener("click",this.confirmButtonListener),this.cancelButton.removeEventListener("click",this.cancelButtonListener)}attributeChangedCallback(e,i,n){switch(e){case"animate":this.animateInOut=n.toLowerCase()==="true";break;case"persist":const l=n.toLowerCase()==="true";if(document.documentElement.removeEventListener("click",this.noPersistHandler),l)return;document.documentElement.addEventListener("click",this.noPersistHandler);break;case"confirm-button-class":this.confirmButton.className=n;break;case"cancel-button-class":this.cancelButton.className=n;break;case"event-data":this.eventData=n;break;case"title":this.titleSlotElement.textContent=n;break;case"message":this.messageSlotElement.textContent=n;break}}static register(){customElements.define("confirm-modal",g,{extends:"dialog"})}}g.register();w.register();const S=document.getElementById("select-all-files"),a=document.getElementById("files-action-form"),r=document.getElementById("files-select-action"),C=document.getElementById("delete-all-selected-files-action"),A=document.querySelectorAll(".file-selected-checkbox"),M=document.getElementById("search-files-button"),c=document.querySelector("confirm-modal"),q=document.querySelectorAll(".media-preview-cell > *"),m=document.querySelector("media-preview-modal");let b=!1;S.addEventListener("change",t=>{const s=t.target,e=document.querySelectorAll(".file-selected-checkbox");b=!0,e.forEach(i=>{i.checked=s.checked}),b=!1,s.checked?r.classList.remove("opacity-0"):r.classList.add("opacity-0")});A.forEach(t=>{t.addEventListener("change",s=>{if(b)return;if(s.target.checked)r.classList.remove("opacity-0");else{const i=document.querySelectorAll(".file-selected-checkbox");let n=!0;i.forEach(l=>{l.checked&&(n=!1)}),n&&r.classList.add("opacity-0")}})});function k(){a.action="/files",a.method="GET",a.submit()}M.addEventListener("click",t=>{t.preventDefault(),k()});document.addEventListener("keyup",t=>{t.key==="Enter"&&k()});a.addEventListener("submit",t=>{t.preventDefault()});C.addEventListener("click",t=>{t.preventDefault();const s=document.querySelectorAll(".file-selected-checkbox:checked").length,e=s>1?"files":"file";c.setAttribute("title","Delete all selected files"),c.setAttribute("message",`Are you sure you want to delete ${s} selected ${e}?`),c.setAttribute("confirm-button-class","button-error"),c.setAttribute("event-data","delete"),c.showModal()});const v=new URL(window.location.href),E=v.searchParams,p=E.get("search");p!=null&&p===""&&(E.delete("search"),v.search=E.toString(),window.history.replaceState({},"",v.toString()));c.addEventListener("cancel",()=>{r.removeAttribute("open")});c.addEventListener("confirm",t=>{r.removeAttribute("open"),t.detail&&t.detail==="delete"&&(a.action="/files/delete",a.method="POST",a.submit())});q.forEach(t=>{const s=t.parentElement;t.addEventListener("click",()=>{var l;const e=s.dataset.type,i=(l=m.shadowRoot)==null?void 0:l.querySelector("#media-preview-slot"),n=s.dataset.url;switch(e){case"image":const u=s.querySelector("img"),d=u.cloneNode(!0);d.slot="media",d.height=u.naturalHeight,d.width=u.naturalWidth,i.replaceChildren(d),m.showModal();break;case"video":const h=document.createElement("video");h.src=n,h.controls=!0,i.replaceChildren(h),m.showModal();break;case"audio":const f=document.createElement("audio");f.src=n,f.controls=!0,i.replaceChildren(f),m.showModal();break}})});