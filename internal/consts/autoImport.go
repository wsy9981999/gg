package consts

const autoImport = `import autoImport from 'unplugin-auto-import/vite'
{{ .importData }}

export default autoImport({
  dts: 'types/auto-imports.d.ts',
  imports:[{{ .importName }}]
  resolvers: [{{ .resolver }}],
  dirs:['src/api/**','src/stores/**','src/hook/**'],
  eslintrc: {
    enabled: true,

  }
})`
