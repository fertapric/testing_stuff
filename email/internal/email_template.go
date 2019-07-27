package internal

const DefaultSubjectTemplate = `
[Elixir CI] Build failed: {{.Repo.FullName}} on {{.CheckSuite.HeadBranch}}
`

const DefauleEmailMarkdownTemplate = `
<p>
  <h2>{{.Event.Repo.FullName}} ({{.CheckSuite.HeadBranch}})</h2>

  <p><a href="{{.Commit.HTMLURL}}">{{.Commit.SHA}}</a> - {{.Commit.Message}}</p>

  <p>The following checks failed:</p>

  <h3>{{ .CheckRun.Name }}</h3>

  <p><b style="color: #cb2431">{{ .CheckRun.Conclusion }}</b> - ran at <b>{{.CheckRun.StartedAt}}</b> in <b>{{.Duration}}</b></p>

  <p><b>{{ .CheckRun.Output.Title }}</b></p>
  <p>{{ .CheckRun.Output.Summary }}</p>
  <p>{{ .CheckRun.Output.Text }}</p>
</p>

<p style="font-size:small;-webkit-text-size-adjust:none;color:#666;">
  &mdash;
  <br />
  You are receiving this because you are the author of the commit.
  <br />
  <a href="{{ .CheckRun.HTMLURL }}">View it on GitHub</a>
</p>
`

const EmailHeaderTemplate = `
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">

<head>
  <!--[if gte mso 9]>
      <xml>
        <o:OfficeDocumentSettings>
        <o:AllowPNG/>
        <o:PixelsPerInch>96</o:PixelsPerInch>
        </o:OfficeDocumentSettings>
      </xml>
    <![endif]-->

  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width">

  <!--[if !mso]><!-->
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <!--<![endif]-->

  <title></title>

  <!--[if !mso]><!-->
  <link href="https://fonts.googleapis.com/css?family=Lato" rel="stylesheet" type="text/css">
  <!--<![endif]-->

  <style type="text/css">
    body {
      margin: 0;
      padding: 0;
    }

    table,
    td,
    tr {
      vertical-align: top;
      border-collapse: collapse;
    }

    * {
      line-height: inherit;
    }

    a[x-apple-data-detectors=true] {
      color: inherit !important;
      text-decoration: none !important;
    }
  </style>

  <style type="text/css" id="media-query">
    @media (max-width: 600px) {

      .block-grid,
      .col {
        min-width: 320px !important;
        max-width: 100% !important;
        display: block !important;
      }

      .block-grid {
        width: 100% !important;
      }

      .col {
        width: 100% !important;
      }

      .col>div {
        margin: 0 auto;
      }

      img.fullwidth,
      img.fullwidthOnMobile {
        max-width: 100% !important;
      }

      .no-stack .col {
        min-width: 0 !important;
        display: table-cell !important;
      }

      .no-stack.two-up .col {
        width: 50% !important;
      }

      .no-stack .col.num4 {
        width: 33% !important;
      }

      .no-stack .col.num8 {
        width: 66% !important;
      }

      .no-stack .col.num4 {
        width: 33% !important;
      }

      .no-stack .col.num3 {
        width: 25% !important;
      }

      .no-stack .col.num6 {
        width: 50% !important;
      }

      .no-stack .col.num9 {
        width: 75% !important;
      }

      .video-block {
        max-width: none !important;
      }

      .mobile_hide {
        min-height: 0px;
        max-height: 0px;
        max-width: 0px;
        display: none;
        overflow: hidden;
        font-size: 0px;
      }

      .desktop_hide {
        display: block !important;
        max-height: none !important;
      }
    }
  </style>
</head>

<body class="clean-body" style="margin: 0; padding: 0; -webkit-text-size-adjust: 100%; background-color: #fff;">

  <!--[if IE]>
    <div class="ie-browser">
  <![endif]-->

  <table class="nl-container" style="table-layout: fixed; vertical-align: top; min-width: 320px; Margin: 0 auto; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #fff; width: 100%;"
    cellpadding="0" cellspacing="0" role="presentation" width="100%" bgcolor="#fff" valign="top">
    <tbody>
      <tr style="vertical-align: top;" valign="top">
        <td style="word-break: break-word; vertical-align: top;" valign="top">
          <!--[if (mso)|(IE)]>
            <table width="100%" cellpadding="0" cellspacing="0" border="0">
              <tr>
                <td align="center" style="background-color:#fff">
          <![endif]-->

          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                <!--[if (mso)|(IE)]>
          <table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:transparent;width:580px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 580px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                      <!--<![endif]-->
                      <div class="mobile_hide">
                        <table class="divider" border="0" cellpadding="0" cellspacing="0" width="100%" style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;"
                          role="presentation" valign="top">
                          <tbody>
                            <tr style="vertical-align: top;" valign="top">
                              <td class="divider_inner" style="word-break: break-word; vertical-align: top; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; padding-top: 5px; padding-right: 5px; padding-bottom: 5px; padding-left: 5px;"
                                valign="top">
                                <table class="divider_content" border="0" cellpadding="0" cellspacing="0" width="100%" style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; border-top: 0px solid transparent; height: 40px;"
                                  align="center" role="presentation" height="40" valign="top">
                                  <tbody>
                                    <tr style="vertical-align: top;" valign="top">
                                      <td style="word-break: break-word; vertical-align: top; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;" height="40" valign="top"><span></span></td>
                                    </tr>
                                  </tbody>
                                </table>
                              </td>
                            </tr>
                          </tbody>
                        </table>
                      </div>
                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
              </div>
            </div>
          </div>
          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #f4effa;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:#f4effa;">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:#f4effa"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:#f4effa;width:580px; border-top: 0px solid #9768d1; border-left: 2px solid #9768d1; border-bottom: 0px solid #9768d1; border-right: 0px solid #9768d1;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 10px; padding-top:0px; padding-bottom:0px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 578px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid #9768d1; border-left:2px solid #9768d1; border-bottom:0px solid #9768d1; border-right:0px solid #9768d1; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 10px;">
                      <!--<![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 5px; padding-left: 5px; padding-top: 10px; padding-bottom: 10px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#d5dae6;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:10px;padding-right:5px;padding-bottom:10px;padding-left:5px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; line-height: 14px; color: #d5dae6;">
                          <p style="font-size: 14px; line-height: 16px; margin: 0;"><strong><span style="font-size: 20px; line-height: 24px; color: #9768d1;">
                          {{.Event.Repo.FullName}}
                          </span></strong></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
              </div>
            </div>
          </div>
          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #fff;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:#fff;">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:#fff"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:#fff;width:580px; border-top: 0px solid transparent; border-left: 2px solid #9768d1; border-bottom: 0px solid transparent; border-right: 0px solid #9768d1;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 10px; padding-top:0px; padding-bottom:0px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 578px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid transparent; border-left:2px solid #9768d1; border-bottom:0px solid transparent; border-right:0px solid #9768d1; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 10px;">
                      <!--<![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 0px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:10px;padding-right:10px;padding-bottom:0px;padding-left:10px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; line-height: 14px; color: #555555;">
                          <p style="font-size: 14px; line-height: 16px; margin: 0;"><span style="font-size: 14px; line-height: 16px; color: #ffffff;">
                          <a href="https://github.com/{{.Event.Repo.FullName}}/tree/{{.CheckSuite.HeadBranch}}" style="color: #9768d1; line-height: 16px; font-size: 14px; text-decoration: none;">
                            <strong>{{.CheckSuite.HeadBranch}}</strong>
                          </a>
                              <span style="color: #808080; line-height: 16px; font-size: 14px;">
                              <span style="line-height: 16px; font-size: 14px;">branch</span>
                              </span></span></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 5px; padding-bottom: 10px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:5px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; line-height: 14px; color: #555555;">
                          <p style="font-size: 14px; line-height: 16px; margin: 0;"><span style="font-size: 14px; line-height: 16px; color: #000000;">

                          <a href="https://github.com/{{.Event.Repo.FullName}}/commit/{{.CheckSuite.HeadSHA}}" style="color: #9768d1; line-height: 16px; font-size: 14px;text-decoration: none;">
                            <strong>
                              <span style="line-height: 16px; font-size: 14px;">
                              {{.SHA}}
                              </span>
                            </strong>
                          </a>
                              <span style="color: #808080; line-height: 16px; font-size: 14px;">
                              by
                              </span>
                              <span style="line-height: 16px; font-size: 14px;">
                                <strong>
                                  <a href="https://github.com/{{.Event.Sender.Login}}" style="color: #808080; line-height: 16px; font-size: 14px; text-decoration: none;">
                                    @{{.Event.Sender.Login}}
                                  </a>
                                </strong>
                              </span>
                              <span style="color: #808080; line-height: 16px; font-size: 14px;">
                              - {{.Commit.Message}}
                              </span>
                            </span></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
              </div>
            </div>
          </div>
`
const EmailCheckTemplate = `
          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #fff;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:#fff;">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:#fff"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:#fff;width:580px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 580px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                      <!--<![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; line-height: 14px; color: #555555;">
                          <p style="font-size: 14px; line-height: 16px; margin: 0;">&nbsp;</p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
              </div>
            </div>
          </div>
          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #f9e9ea;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:#f9e9ea;">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:#f9e9ea"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:#f9e9ea;width:580px; border-top: 0px solid transparent; border-left: 2px solid #cb2431; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top:5px; padding-bottom:5px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 578px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid transparent; border-left:2px solid #cb2431; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 10px; padding-left: 10px;">
                      <!--<![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 5px; padding-left: 5px; padding-top: 5px; padding-bottom: 5px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#cb2431;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:5px;padding-right:5px;padding-bottom:5px;padding-left:5px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; line-height: 14px; font-size: 12px; color: #cb2431;">
                          <p style="line-height: 14px; font-size: 12px; margin: 0;">
                            <strong>
                            <a href="{{.CheckRun.HTMLURL}}" style="font-size: 18px; line-height: 21px; color: #cb2431; text-decoration: none;">
                              {{ .CheckRun.Name }}
                            </a>
                            </strong></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
              </div>
            </div>
          </div>
          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #fff;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:#fff;">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:#fff"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:#fff;width:580px; border-top: 0px solid #cb2431; border-left: 2px solid #cb2431; border-bottom: 0px solid #cb2431; border-right: 0px solid #cb2431;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top:10px; padding-bottom:10px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 578px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid #cb2431; border-left:2px solid #cb2431; border-bottom:0px solid #cb2431; border-right:0px solid #cb2431; padding-top:10px; padding-bottom:10px; padding-right: 10px; padding-left: 10px;">
                      <!--<![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 5px; padding-left: 5px; padding-top: 5px; padding-bottom: 5px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:5px;padding-right:5px;padding-bottom:5px;padding-left:5px;">
                        <div style="font-size: 12px; line-height: 14px; font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; color: #555555;">
                          <p style="font-size: 14px; line-height: 14px; margin: 0;"><span style="font-size: 12px;"><span style="color: #808080; line-height: 14px; font-size: 12px;">ran at</span> <span style="color: #000000; line-height: 14px; font-size: 12px;"><strong>
                          {{.CheckRun.StartedAt}}
                          </strong></span> <span style="color: #808080; line-height: 14px; font-size: 12px;">in</span> <span style="color: #000000; line-height: 14px; font-size: 12px;"><strong>
                          {{.Duration}}
                          </strong></span></span></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 5px; padding-left: 5px; padding-top: 10px; padding-bottom: 5px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#132F40;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:10px;padding-right:5px;padding-bottom:5px;padding-left:5px;">
                        <div style="line-height: 14px; font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; color: #132F40;">
                          <p style="line-height: 19px; font-size: 12px; margin: 0;">
                          <span style="font-size: 16px; color: #000000;">
                          {{ .CheckRun.Output.Summary }}
                          </span></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 5px; padding-left: 5px; padding-top: 15px; padding-bottom: 5px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:15px;padding-right:5px;padding-bottom:5px;padding-left:5px;">
                        <div style="font-size: 12px; line-height: 14px; font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; color: #555555;">
                          <p style="font-size: 14px; line-height: 14px; margin: 0;">
                          <span style="font-size: 12px;">
                          <span style="color: #000000; line-height: 14px; font-size: 12px;">
                          <strong>
                          <a href="{{.CheckRun.HTMLURL}}" style="color: #cb2431; line-height: 14px; font-size: 12px;">
                            <span style="line-height: 14px; font-size: 12px;">
                              View on GitHub
                            </span>
                          </a>
                          </strong>
                            </span></span></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->

                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>

                <!--[if (mso)|(IE)]>
  </td>
  </tr>
  </table>
  <![endif]-->

                <!--[if (mso)|(IE)]>
  </td>
  </tr>
  </table>
  </td>
  </tr>
  </table>

  <![endif]-->

              </div>
            </div>
          </div>
`

const EmailFooterTemplate = `
          <div style="background-color:transparent;">
            <div class="block-grid " style="Margin: 0 auto; min-width: 320px; max-width: 580px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #fff;">
              <div style="border-collapse: collapse;display: table;width: 100%;background-color:#fff;">
                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:580px"><tr class="layout-full-width" style="background-color:#fff"><![endif]-->
                <!--[if (mso)|(IE)]><td align="center" width="580" style="background-color:#fff;width:580px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px;"><![endif]-->
                <div class="col num12" style="min-width: 320px; max-width: 580px; display: table-cell; vertical-align: top; width: 580px;">
                  <div style="width:100% !important;">
                    <!--[if (!mso)&(!IE)]><!-->
                    <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                      <!--<![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; line-height: 14px; color: #555555;">
                          <p style="font-size: 14px; line-height: 16px; margin: 0;">&nbsp;</p>
                        </div>
                      </div>

                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px; font-family: Tahoma, Verdana, sans-serif"><![endif]-->
                      <div style="color:#555555;font-family:'Lato', Tahoma, Verdana, Segoe, sans-serif;line-height:120%;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                        <div style="font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; font-size: 12px; line-height: 14px; color: #555555;">
                          <p style="font-size: 14px; line-height: 14px; text-align: center; margin: 0;"><span style="font-size: 12px; color: #808080;">
                          You are receiving this email because your are {{.WhyReceiving}}</span></p>
                        </div>
                      </div>
                      <!--[if mso]></td></tr></table><![endif]-->
                      <!--[if (!mso)&(!IE)]><!-->
                    </div>
                    <!--<![endif]-->
                  </div>
                </div>
                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
              </div>
            </div>
          </div>
          <!--[if (mso)|(IE)]>
                  </td>
                </tr>
              </table>
            <![endif]-->
        </td>
      </tr>
    </tbody>
  </table>

  <!--[if (IE)]>
      </div>
    <![endif]-->
</body>

</html>
`
