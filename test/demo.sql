SELECT
    lis.patient_id AS "patientId",
    '' AS "series",
    '' AS "admissionId",
    lis.report_id AS "reportId",
    lis.report_name AS "reportName",
    to_char (
        lis.report_date, 'yyyy-MM-dd HH24:mm:ss'
    ) AS "reportDate",
    lis.reporter_name as "reporterName",
    lis.reporter_code as "reporterCode",
    lisdetail.item_id as "itemId",
    lisdetail.item_name as "itemName",
    lisdetail.report_details as "reportDetails",
    lisdetail.report_abnormal_sign as "reportAbnormalSign",
    lisdetail.high_max_value as "highMaxValue",
    lisdetail.low_max_value as "lowMaxValue",
    lisdetail.remark as "remark"
FROM
    v_yyt_mz_lis lis
    inner join v_yyt_mz_lis_detail lisdetail on lis.report_id = lisdetail.report_id
WHERE
    1 != 1 < if test = 'patientId !=null and patientId !=""' >
    OR lis.patient_id = #{patientId}
    < / if > < if test = "StartDate != null and StartDate != ''" >
    AND lis.report_date & gt;

=to_date(#{StartDate} || ' 00:00:00','yyyy-mm-dd hh24:mi:ss')
  </if>
  	<if test="EndDate != null and EndDate != ''">
    AND lis.report_date&lt;

=to_date(#{EndDate} || ' 23:59:59','yyyy-mm-dd hh24:mi:ss')
  </if>
