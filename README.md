# zeng - Zendesk Guide Command Line Tool (Experimental)

Zeng is a CLI to manage an article on Zendesk Guide.
This tool is still experimental. Commands on CLI might be changed dramatically.

# Installation

## Mac

TBD

## Windows

TBD

# How To Use

## Configuration

### Email / Password
First, run the following command and input your credentials, then zeng stores the configuration to ~/.zeng/zeng.conf.

```
$ zeng configure
```

```
$ zeng configure
Using config file: /Users/takahashi/.zeng.yaml
What is Zendesk Subdomain (subdomain.zendesk.com)?
Enter a value: yourcompany

What is your Zendesk Email?
Enter a value: sample@sample.com

What is your default locale?
Enter a value (Default is en-us):

How do you prefer to acccess to Zendesk?

1. password
2. apikey

Enter a number (Default is 1): 1

What is your Zendesk password?
Enter a value: **************
```

Limitation: If you enable 2-factor authentication, you need to use the following APIKEY instead of Email/Password. Because 2-factor authentication doesn't allow Email/Password login.

### APIKEY

```
$ zeng configure
Using config file: /Users/takahashi/.zeng.yaml
What is Zendesk Subdomain (subdomain.zendesk.com)?
Enter a value: yourcompany

What is your Zendesk Email?
Enter a value: sample@sample.com

What is your default locale?
Enter a value (Default is en-us):

How do you prefer to acccess to Zendesk?

1. password
2. apikey

Enter a number (Default is 1): 2

What is your Zendesk APIKEY?
Enter a value: ****************************************
```


### OAuth

TODO - NOT Supported

## Sample Use-Case

### Download an article (WIP)

Get Article command exports a speficied article to the following structure.

```
$ zeng get article <article_id>
```

```
a_<article id>/
    a_<article_id>_<locale>_<title>.html
    meta_<article_id>_<locale>.yaml
```

meta_<atricle_id>.yaml contains the following information as same as https://developer.zendesk.com/rest_api/docs/help_center/articles excluding `body` parameter.

```
```

### Donwload articles inside of a section (WIP)

Get Section command exports articles in a speficied section to the following structure.

```
$ zeng get section <senction_id>
```

```
s_<section_id>_<locale>/
    a_<article id>/
        a_<article_id>_<locale>_<title>.html
        meta_<article_id>_<locale>.yaml
    a_<article id>/
        a_<article_id>_<locale>_<title>.html
        meta_<article_id>.yaml
```

### Donwload articles inside of a category (WIP)

Get Category command exports articles in a speficied category to the following structure.

```
$ zeng get category <category_id>
```

```
c_<category_id>/
    s_<section_id>/
        a_<article id>/
            a_<article_id>_<locale>_<title>.html
            meta_a_<article_id>_<locale>.yaml
    s_<section_id>/
        a_<article id>/
            a_<article_id>_<locale>_<title>.html
            meta_<article_id>_<locale>.yaml
```

### Donwload all article (WIP)

Get Guide command exports all article.

```
$ zeng get guide
```

```
zendesk_<subdomain>/
  c_<category_id>_<locale>.
      s_<section_id>_<locale>/
          a_<article id>_<locale>/
              a_<article_id>_<title>.html
              meta_<article_id>.yaml
      s_<section_id>_<locale>/
          a_<article id>_<locale>/
              a_<article_id>_<title>.html
              meta_<article_id>.yaml
```

### Preview an updated article

Open a specified article by using your default browser quickly.

```
$ zeng preview <article_id>
```

### List articles

list commands outputs a list of articles with a table format

```
$ zeng list article
Using config file: /Users/xxxx/.zeng.yaml
Info: Request to https://xxxxx.zendesk.com/api/v2/help_center/en-us/articles.json?include=categories%2Csections&page=1&per_page=20&sort_by=position&sort_order=asc
+--------------+--------------------------------+--------------+--------------------------------+--------------+--------------------------------+------------------------------------------------------------------------------------------------------------------------------+----------------------+
| CATEGORY ID  |            CATEGORY            |  SECTION ID  |            SECTION             |  ARTICLE ID  |             TITLE              |                                                             URL                                  |      EDITED AT       |
+--------------+--------------------------------+--------------+--------------------------------+--------------+--------------------------------+------------------------------------------------------------------------------------------------------------------------------+----------------------+
| 000000001 | Privacy, Security and          | 000000001 | Permissions (Alpha)            | 000000001 | *Access Control (Alpha)        | https://support.treasuredata.com/hc/en-us/articles/000000001--Access-Control-Alpha-                                  | 2018-08-09T19:08:51Z | |              | Administration                 |              |                                |              |                                |                                  |                      |
| 0000000012 | Privacy, Security and          | 000000001 | Accounts (Alpha)               | 000000001 | Get API Keys (Alpha)           | https://support.treasuredata.com/hc/en-us/articles/0000000012-Get-API-Keys-Alpha-                                  | 2018-08-09T23:18:53Z | |              | Administration                 |              |                                |              |                                |                                  |                      |
```

### List sections

```
$ zeng list section
$ zeng list section -p 2
$ zeng list section -p 1 -l 100
```

### List categories

```
$ zeng list category
```

### TBD - Search Articles
### TBD - Update an existing article
### TBD - Create an new article
### TBD - Archive an existing article