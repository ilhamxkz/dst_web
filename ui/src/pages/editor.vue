<template>
  <div class="main-container">
    <!-- Controls -->
    <div class="template-controls" style="max-width:100vw; margin:0 auto 10px auto; padding:0 10px;">
      <div style="margin-bottom:12px;">
        <label style="display:block; margin-bottom:8px; font-weight:600; color:#374151;">
          Choose Template:
        </label>
        <select v-model="selectedTemplate" style="padding:8px 12px; border:1px solid #d1d5db; border-radius:6px; min-width:220px;">
          <option value="">-- Select Template --</option>
          <option v-for="t in availableTemplates" :key="t.name" :value="t.name">{{ t.title }}</option>
        </select>
      </div>

      <div style="display:flex; gap:12px; align-items:center; flex-wrap:wrap;">
        <button @click="loadSelectedTemplate" :disabled="loading || !selectedTemplate" class="btn btn-green">
          {{ loading ? "Loading..." : "Load Template" }}
        </button>

        <button @click="loadAllTemplates" :disabled="loadingTemplates" class="btn btn-blue">
          {{ loadingTemplates ? "Loading..." : "Refresh Templates" }}
        </button>

        <button @click="clearEditor" class="btn btn-red">Clear Editor</button>

        <span style="color:#6b7280; font-size:14px; margin-left:10px;">{{ templateStatus }}</span>
      </div>
    </div>

    <!-- CKEditor element (we init editor on this div) -->
    <div ref="editorContainer" style="max-width:75vw;">
      <div ref="editorElement" class="editor-placeholder"></div>
    </div>

    <!-- Save -->
    <button @click="savePost" :disabled="saving" class="save-button">
      {{ saving ? "Saving..." : "Save" }}
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from "vue";
// Make sure you installed: npm i @ckeditor/ckeditor5-build-classic
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";

/**
 * Template Editor component
 *
 * Behavior:
 * - loadAllTemplates() => GET /api/templates
 * - loadSelectedTemplate() => GET /api/templates/:name
 *    -> extracts <style>..</style> blocks and inject to <head> as <style id="template-styles-{name}">
 *    -> sets editor data with HTML WITHOUT those <style> tags
 * - clearEditor() => clear editor and remove injected style for selected template
 * - savePost() => compose finalContent = (injectedStyles ? `<style>..</style>` : '') + editorData
 *               => POST /api/posts
 */

// reactive state
const availableTemplates = ref([]);
const selectedTemplate = ref("");
const templateStatus = ref("Ready to load template");
const loadingTemplates = ref(false);
const loading = ref(false);
const saving = ref(false);

// editor refs & instance
const editorElement = ref(null);
const editorContainer = ref(null);
let editorInstance = null;

// Helper: extract <style> blocks, inject into head (id = template-styles-{name}), return htmlWithoutStyles
function extractAndInjectStyles(html = "", name = "template") {
  const styleRegex = /<style[^>]*>[\s\S]*?<\/style>/gi;
  const matches = html.match(styleRegex) || [];
  const combinedStyles = matches.join("\n");
  const htmlWithoutStyles = html.replace(styleRegex, "");

  // inject into head with unique id
  const styleId = `template-styles-${name}`;
  // remove any existing with same id first
  const existing = document.getElementById(styleId);
  if (existing) existing.remove();

  if (combinedStyles.trim().length > 0) {
    const styleEl = document.createElement("style");
    styleEl.id = styleId;
    styleEl.dataset.origin = name;
    styleEl.innerHTML = combinedStyles;
    document.head.appendChild(styleEl);
  }

  return { htmlWithoutStyles, combinedStyles };
}

// Helper: remove injected style by template name
function removeInjectedStyles(name = "template") {
  const styleId = `template-styles-${name}`;
  const el = document.getElementById(styleId);
  if (el) el.remove();
}

// Load list of templates
async function loadAllTemplates() {
  loadingTemplates.value = true;
  templateStatus.value = "Loading available templates...";
  try {
    const resp = await fetch("http://localhost:8080/api/templates");
    if (!resp.ok) throw new Error(`HTTP ${resp.status}`);
    const data = await resp.json();
    availableTemplates.value = data.templates || [];
    templateStatus.value = `${availableTemplates.value.length} templates available`;
  } catch (err) {
    console.error(err);
    templateStatus.value = "Failed to load templates: " + err.message;
    alert("Gagal load daftar template: " + err.message);
  } finally {
    loadingTemplates.value = false;
    setTimeout(() => (templateStatus.value = "Ready to load template"), 1500);
  }
}

// Load selected template, extract & inject styles, set editor data (without style tags)
async function loadSelectedTemplate() {
  if (!selectedTemplate.value) {
    alert("Pilih template terlebih dahulu");
    return;
  }
  loading.value = true;
  templateStatus.value = `Loading ${selectedTemplate.value} template...`;

  try {
    const resp = await fetch(`http://localhost:8080/api/templates/${selectedTemplate.value}`);
    if (!resp.ok) {
      const txt = await resp.text();
      throw new Error(`HTTP ${resp.status}: ${txt}`);
    }
    const data = await resp.json();
    const htmlRaw = data.html || "";

    // remove previously injected styles for any other template
    // (if you want to keep multiple, change logic)
    availableTemplates.value.forEach((t) => {
      if (t.name && t.name !== selectedTemplate.value) removeInjectedStyles(t.name);
    });

    // extract & inject styles
    const { htmlWithoutStyles } = extractAndInjectStyles(htmlRaw, selectedTemplate.value);

    // set editor content (use setData if editor ready)
    if (editorInstance && typeof editorInstance.setData === "function") {
      editorInstance.setData(htmlWithoutStyles);
    } else {
      // fallback store to editor later (shouldn't happen)
      console.warn("Editor not ready yet, storing initial content.");
      // create a small delay and retry setData when editor initializes
      setTimeout(() => editorInstance && editorInstance.setData && editorInstance.setData(htmlWithoutStyles), 300);
    }

    templateStatus.value = `${selectedTemplate.value} template loaded successfully!`;
  } catch (err) {
    console.error(err);
    templateStatus.value = "Failed to load template: " + err.message;
    alert("Gagal load template: " + err.message);
  } finally {
    loading.value = false;
    setTimeout(() => (templateStatus.value = "Ready to load template"), 1500);
  }
}

// Clear editor & remove associated injected style
function clearEditor() {
  // clear editor content
  if (editorInstance && typeof editorInstance.setData === "function") {
    editorInstance.setData("");
  }
  // remove injected style for selected template (if any)
  if (selectedTemplate.value) removeInjectedStyles(selectedTemplate.value);
  selectedTemplate.value = "";
  templateStatus.value = "Editor cleared";
  setTimeout(() => (templateStatus.value = "Ready to load template"), 1200);
}

// Save post: we will gather injected CSS (if exists) and editor HTML and send to API
async function savePost() {
  if (!editorInstance) {
    alert("Editor belum siap.");
    return;
  }

  try {
    saving.value = true;

    // get editor HTML
    const editorHtml = editorInstance.getData() || "";

    // get injected style for current template if exists
    let combinedStyles = "";
    if (selectedTemplate.value) {
      const styleEl = document.getElementById(`template-styles-${selectedTemplate.value}`);
      if (styleEl) combinedStyles = styleEl.innerHTML || "";
    }

    // compose final content: include style tag before body HTML
    const finalContent = (combinedStyles && combinedStyles.trim().length > 0)
      ? `<style>\n${combinedStyles}\n</style>\n` + editorHtml
      : editorHtml;

    const payload = {
      title: `Edited ${selectedTemplate.value || "Custom"} Template - ${new Date().toLocaleString()}`,
      content: finalContent,
      author_id: null,
      template_name: selectedTemplate.value || "custom"
    };

    const resp = await fetch("http://localhost:8080/api/posts", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload)
    });

    if (!resp.ok) {
      const txt = await resp.text();
      throw new Error(`HTTP ${resp.status}: ${txt}`);
    }

    const data = await resp.json();
    alert(`Template berhasil disimpan!\nID: ${data.post?.id || data.post_id || "unknown"}\nTitle: ${data.post?.title || payload.title}`);
    templateStatus.value = "Template saved successfully!";
    setTimeout(() => (templateStatus.value = "Ready to load template"), 1500);
  } catch (err) {
    console.error(err);
    alert("Gagal menyimpan: " + err.message);
  } finally {
    saving.value = false;
  }
}

// Initialize CKEditor on mount
onMounted(async () => {
  await loadAllTemplates();

  // create editor instance
  try {
    // basic config. adjust toolbar/plugins as needed for your build
    const config = {
      toolbar: {
        items: [
          "undo", "redo", "|",
          "bold", "italic", "underline", "link", "|",
          "insertTable", "bulletedList", "numberedList", "|",
          "blockQuote", "removeFormat"
        ],
        shouldNotGroupWhenFull: true
      },
      placeholder: "Select a template above or start typing...",
      table: { contentToolbar: ["tableColumn", "tableRow", "mergeTableCells"] }
    };

    editorInstance = await ClassicEditor.create(editorElement.value, config);

    // when content changes, you can react (if you want to keep a reactive var)
    editorInstance.model.document.on("change:data", () => {
      // optional: you can access editorInstance.getData() if needed
      // console.log("Editor changed. length:", editorInstance.getData().length);
    });
  } catch (err) {
    console.error("Failed to initialize editor:", err);
    alert("Editor initialization failed: " + (err.message || err));
  }
});

// Destroy editor on unmount
onBeforeUnmount(() => {
  if (editorInstance) {
    editorInstance.destroy().catch((err) => console.warn("Error destroying editor:", err));
    editorInstance = null;
  }
});
</script>

<style scoped>
.main-container {
  width: auto;
  height: auto;
}

/* simple button classes */
.btn { padding:8px 14px; border-radius:6px; color:white; border:none; cursor:pointer; font-size:14px; }
.btn-green { background:#16a34a; }
.btn-blue { background:#2563eb; }
.btn-red { background:#dc2626; }

/* Editor placeholder styling */
.editor-placeholder {
  border: 1px solid #e5e7eb;
  min-height: 620px; 
  border-radius: 8px;
  padding: 0px;
  background: white;
  box-shadow: 0 2px 10px rgba(0,0,0,0.03);
  width: 100%;
}


/* small responsive tweaks */
@media (max-width: 768px) {
  .editor-placeholder { min-height: 320px; }
  .save-button { right: 12px; bottom: 12px; padding: 10px 14px; }
}

/* Tombol simpan (fixed) tetap terlihat */
.save-button {
  position: fixed;
  right: 24px;
  bottom: 24px;
  background: #2563eb;
  color: white;
  border: none;
  padding: 12px 18px;
  border-radius: 10px;
  box-shadow: 0 6px 18px rgba(37,99,235,0.15);
  cursor: pointer;
  z-index: 9999;
  font-size: 14px;
  min-width: 88px;
}

/* Responsive: di layar sempit, jangan kalkulasi sidebar */
@media (max-width: 1100px) {
  .editor-wrap,
  .template-controls,
  .template-controls > div {
    width: 95vw !important;
  }
  .editor-placeholder { min-height: 420px; }
  .save-button { right: 12px; bottom: 12px; padding: 10px 14px; }
}
</style>
