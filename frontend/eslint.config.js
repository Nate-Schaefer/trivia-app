import js from '@eslint/js'                                                                                                                                                                                        
import react from 'eslint-plugin-react'                                                                                                                                                                            
import reactHooks from 'eslint-plugin-react-hooks'    
import globals from 'globals'                                                                                                                                                             
                                                                                                                                                                                                                    
export default [                                                                                                                                                                                                   
    js.configs.recommended,
    react.configs.flat.recommended,                                                                                                                                                                                  
    react.configs.flat['jsx-runtime'],
    {                                                                                                                                                                                                                
        files: ['src/**/*.{js,jsx}'],
        plugins: {
        'react-hooks': reactHooks,
        },
        languageOptions: {
            globals: {
                ...globals.browser,
                ...globals.jest,
            },
            parserOptions: {                                                                                                                                                                                             
                ecmaFeatures: {
                   jsx: true,                                                                                                                                                                                               
                },      
            },
        },
        settings: {                                                                                                                                                                                                        
            react: {                                                                                                                                                                                                         
              version: '19',                                                                                                                                                                                             
            },                                                                                                                                                                                                               
        },
        rules: {
            ...reactHooks.configs.recommended.rules,
        },
    },
]