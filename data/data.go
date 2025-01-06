package data

var DefaultJson string = `
[
            {
                "id": "header",
                "type": "section",
                "style": {
                    "display": "flex",
                    "justifyContent": "space-between",
                    "alignItems": "center",
                    "padding": "20px",
                    "backgroundColor": "#ffffff",
                    "width": "100%",
                    "height": "5rem"
                },
                "children": [
                    {
                        "id": "logo",
                        "type": "text",
                        "content": "Your Hackathon",
                        "style": {
                            "fontSize": "20px",
                            "fontWeight": "bold",
                            "color": "#333333"
                        }
                    },
                    {
                        "id": "nav-section",
                        "type": "section",
                        "style": {
                            "display": "flex",
                            "gap": "10px",
                            "alignItems": "center"
                        },
                        "children": [
                            { "id": "link-1", "type": "link", "content": "Home", "style": { "color": "#333", "textDecoration": "none" } },
                            { "id": "link-2", "type": "link", "content": "Rules", "style": { "color": "#333", "textDecoration": "none" } },
                            { "id": "link-3", "type": "link", "content": "Gallery", "style": { "color": "#333", "textDecoration": "none" } },
                            { "id": "link-4", "type": "link", "content": "Contact Us", "style": { "color": "#333", "textDecoration": "none" } }
                        ]
                    }
                ]
            },
            {
                "id": "hero-section",
                "type": "section",
                "style": {
                    "display": "flex",
                    "align-items": "center",
                    "flex-direction": "column",
                    "justify-content": "center",
                    "height": "80vh"
                },
                "children": [
                    {
                        "id": "hero-section",
                        "type": "section",
                        "style": {
                            "backgroundColor": "#e8d5f7",
                            "borderRadius": "10px",
                            "margin": "20px",
                            "height": "50vh",
                            "width": "90%",
                            "display": "flex",
                            "align-items": "center",
                            "justify-content": "center",
                            "flex-direction": "column"
                        },
                        "children": [
                            {
                                "id": "title",
                                "type": "text",
                                "content": "Your Hackathon Name",
                                "style": { "fontSize": "30px", "fontWeight": "bold", "color": "#333" }
                            }
                        ]
                    },
                    {
                        "id": "dates",
                        "type": "section",
                        "style": {
                            "display": "flex",
                            "justifyContent": "space-around",
                            "height": "10rem",
                            "background-color": "black",
                            "margin": "20px",
                            "borderRadius": "10px",
                            "position": "relative",
                            "bottom": "5rem",
                            "width": "80%",
                            "align-items": "center"
                        },
                        "children": [
                            {
                                "id": "start-date",
                                "type": "text",
                                "content": "2023/12/01",
                                "style": { "fontSize": "16px", "fontWeight": "500", "color": "white" }
                            },
                            {
                                "id": "end-date",
                                "type": "text",
                                "content": "2023/12/10",
                                "style": { "fontSize": "16px", "fontWeight": "500", "color": "white" }
                            }
                        ]
                    }
                ]
            },
            {
                "id": "hero-section",
                "type": "section",
                "style": {
                    "display": "flex",
                    "align-items": "center",
                    "flex-direction": "column",
                    "justify-content": "center",
                    "height": "85vh",
                    "width":"100%"
                },
                "children": [
                    {
                        "id": "rounds-section",
                        "type": "section",
                        "style": {
                            "display": "flex",
                            "align-items": "flex-start",
                            "flex-direction": "column",
                            "justify-content": "flex-start",
                            "height": "85vh",
                            "padding": "40px",
                            "width":"100%"
                        },
                        "children": [
                            {
                                "id": "rules-title",
                                "type": "text",
                                "content": "Rules And Rounds",
                                "style": { "fontSize": "24px", "fontWeight": "bold", "color": "#333" }
                            },
                            {
                                "id": "rules-content",
                                "type": "section",
                                "style": {
                                    "display": "grid",
                                    "gridTemplateColumns": "1fr",
                                    "gap": "20px",
                                    "marginTop": "20px",
                                    "width":"100%"
                                },
                                "children": [
                                    {
                                        "id": "start-date",
                                        "type": "text",
                                        "style": { "fontSize": "16px", "fontWeight": "500", "color": "black", "width":"100%"},
                                        "content": "elksd rtkflmd krjkfld rjf  jrnf jfkdm trkjrn rkje,ndswjk jj4 jrktngrk gk bjrt,wnegj4kw rwken, gjk, gr 4tkjer, gjk4 er.j kj"
                                    }
                                ]
                            },
                            {
                                "id": "rules-content",
                                "type": "section",
                                "style": {
                                    "display": "flex",
                                    "flex-direction": "row",
                                    "align-items":"flex-start",
                                    "width": "100%",
                                    "height": "100%",
                                    "padding":"1rem"
                                },
                                "children": [
                                    {
                                        "id": "rules-listing",
                                        "type": "section",
                                        "style": {
                                            "width": "50%",
                                            "height": "100%",
                                            "display":"flex",
                                            "align-items":"center",
                                            "justify-content":"center",
                                            "flex-direction":"column"
                                        },
                                        "children": [
                                            {
                                                "id": "rules",
                                                "type": "text",
                                                "style": {
                                                    "width": "100%",
                                                    "border-radius": "0.5rem",
                                                    "padding": "1rem",
                                                    "margin": "1rem 0",
                                                    "height": "fit-content",
                                                    "border": "1px solid black"
                                                },
                                                "content": "rule1"
                                            }, {
                                                "id": "rules",
                                                "type": "text",
                                                "style": {
                                                    "width": "100%",
                                                    "border-radius": "0.5rem",
                                                    "padding": "1rem",
                                                    "margin": "1rem 0",
                                                    "height": "fit-content",
                                                    "border": "1px solid black"
                                                },
                                                "content": "rule2"
                                            }, {
                                                "id": "rules",
                                                "type": "text",
                                                "style": {
                                                    "width": "100%",
                                                    "border-radius": "0.5rem",
                                                    "padding": "1rem",
                                                    "margin": "1rem 0",
                                                    "height": "fit-content",
                                                    "border": "1px solid black"
                                                },
                                                "content": "rule3"
                                            }
                                        ]
                                    },
                                    {
                                        "id": "rules-description",
                                        "type": "section",
                                        "style": {
                                            "width": "50%",
                                            "height": "100%",
                                            "padding":"25px"
                                        },
                                        "children": [
                                            {
                                                "id": "text",
                                                "type": "section",
                                                "style": {
                                                    "padding":"20px",
                                                    "width":"100%",
                                                    "height":"100%",
                                                    "display":"flex",
                                                    "align-items":"center",
                                                    "justify-content":"center",
                                                    "background-color": "#e8d5f7",
                                                    "box-shadow":"15px 20px rgba(1, 1, 1, 1)",
                                                    "border-radius":"1rem"
                                                },
                                                "children":[
                                                    {
                                                        "id":"text",
                                                        "type":"section",
                                                        "style":{},
                                                        "children":[
                                                            {
                                                                "id":"text",
                                                                "type":"text",
                                                                "style":{},
                                                                "content": "klfmgolg bkjgdm jg hrv nm njhgf ehjtnbjhknb njrhbk hjrbhj rgtke jt"
                                                            }
                                                        ]
                                                    }
                                                ]
                                                
                                            }
                                        ]
                                    }
                                ]

                            }
                        ]
                    }
                ]
            },
            {
                "id": "registration-section",
                "type": "section",
                "style": {
                    "height":"fit-content",
                    "padding": "20px",
                    "margin": "20px"
                },
                "children": [
                    {
                        "id":"registeration-button-container",
                        "type":"section",
                        "style":{
                            "display":"flex",
                            "align-items":"center",
                            "flex-direction":"column",
                            "height":"100%"
                        },
                        "children":[
                            {
                                "id": "registration-title",
                                "type":"button",
                                "content":"Get Registered",
                                "style":{
                                    "outline":"none",
                                    "border":"2px black solid",
                                    "background":"none",
                                    "width":"15%",
                                    "padding":"1rem",
                                    "border-radius":"2rem",
                                    "position":"relative",
                                    "top":"1.5rem",
                                    "background-color":"white",
                                    "cursor":"pointer"
                                }
                            },
                            {
                                "id": "registration-details",
                                "type": "section",
                                "style":{
                                    "width":"90%",
                                    "height":"fit-content",
                                    "border":"1px black solid",
                                    "text-align":"center"
                                },
                                "children":[
                                    {
                                        "id":"registeration-container",
                                        "type":"text",
                                        "content": "Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. v Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here. Provide the registration details, links, or forms here.",
                                        "style": { 
                                            "fontSize": "16px", 
                                            "color": "#333",
                                            "padding":"3rem"
                                        }
                                    }
                                ]
                                
                            }
                        ]
                    }
                ]
            },
            {
                "id":"gallery",
                "type":"section",
                "style":{
                    "display":"flex",
                    "height":"fit-content",
                    "padding":"0 2rem",
                    "flex-direction":"column"
                },
                "children":[
                    {
                        "id":"intro-gallery",
                        "type":"section",
                        "style":{
                            "width":"100%",
                            "height":"fit-content"
                        },
                        "children":[
                            {
                                "id": "gallery-title",
                                "type": "text",
                                "content": "Gallery",
                                "style": { "fontSize": "24px", "fontWeight": "bold", "color": "#333", "text-align":"center" }
                            }
                        ]
                    },
                    {
                        "id":"photos",
                        "type":"section",
                        "style":{
                            "display":"flex",
                            "align-items":"center",
                            "justify-content":"center",
                            "flex-wrap":"wrap",
                            "gap":"5rem",
                            "padding":"5rem"
                        },
                        "children":[
                            {
                                "id":"photo1",
                                "type":"section",
                                "style":{
                                    "height":"20rem",
                                    "width":"20rem",
                                    "background-position":"center",
                                    "background-size":"cover",
                                    "background-image":"url('https://images.pexels.com/photos/3913031/pexels-photo-3913031.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1')"
                                },
                                "children":[]
                            },{
                                "id":"photo1",
                                "type":"section",
                                "style":{
                                    "height":"20rem",
                                    "width":"20rem",
                                    "background-position":"center",
                                    "background-size":"cover",
                                    "background-image":"url('https://images.pexels.com/photos/3913031/pexels-photo-3913031.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1')"
                                },
                                "children":[]
                            },{
                                "id":"photo1",
                                "type":"section",
                                "style":{
                                    "height":"20rem",
                                    "width":"20rem",
                                    "background-position":"center",
                                    "background-size":"cover",
                                    "background-image":"url('https://images.pexels.com/photos/3913031/pexels-photo-3913031.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1')"
                                },
                                "children":[]
                            }
                        ]
                    }
                ]
            },
            {
                "id": "contact",
                "type": "section",
                "style": {
                    "padding": "20px",
                    "color": "#ffffff",
                    "height":"fit-content"
                },
                "children": [
                    {
                        "id": "contact-text",
                        "type": "text",
                        "content": "Contact Us",
                        "style": { "fontSize": "24px", "fontWeight": "bold", "color": "#333","padding":"2rem 0"}
                    },
                    {
                        "id": "contact-text",
                        "type": "text",
                        "content": "Contact Member Name",
                        "style": { "fontSize": "20px","fontWeight": "500", "color": "black","padding-top":"1rem"}
                    },
                    {
                        "id": "contact-text",
                        "type": "text",
                        "content": "987654321",
                        "style": { "fontSize": "15px","fontWeight": "400", "color": "black","padding-top":"0.1rem"}
                    },
                    {
                        "id": "contact-text",
                        "type": "text",
                        "content": "Contact Member Name",
                        "style": { "fontSize": "20px","fontWeight": "500", "color": "black","padding-top":"1rem"}
                    },
                    {
                        "id": "contact-text",
                        "type": "text",
                        "content": "987654321",
                        "style": { "fontSize": "15px","fontWeight": "400",  "color": "black", "padding":"0.1rem 0"}
                    }
                ]
            },
            {
                "id": "footer",
                "type": "section",
                "style": {
                    "padding": "20px",
                    "backgroundColor": "#333333",
                    "color": "#ffffff",
                    "textAlign": "center"
                },
                "children": [
                    {
                        "id": "footer-text",
                        "type": "text",
                        "content": "© 2024 Your Hackathon. All rights reserved.",
                        "style": { "fontSize": "14px", "color": "#ffffff" }
                    }
                ]
            }
        ]`
