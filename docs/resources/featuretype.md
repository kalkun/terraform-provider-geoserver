---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "geoserver_featuretype Resource - terraform-provider-geoserver"
subcategory: ""
description: |-
  
---

# geoserver_featuretype (Resource)



## Example Usage

```terraform
resource "geoserver_featuretype" "bd_topo_batiment" {
  workspace_name = geoserver_workspace.my_workspace.name
  datastore_name = geoserver_datastore.db-referentiels-carto-bd_topo-ign.name
  name           = "fdp_batiment"
  native_name    = "batiment"
  enabled        = true

  lat_lon_bounding_box_max_x     = 55.83018112182617
  lat_lon_bounding_box_max_y     = 51.08795166015625
  lat_lon_bounding_box_min_x     = -63.152530670166016
  lat_lon_bounding_box_min_y     = -21.387481689453125
  lat_lon_bounding_box_crs_value = "EPSG:4326"

  native_bounding_box_max_x     = 55.83018112182617
  native_bounding_box_max_y     = 51.08795166015625
  native_bounding_box_min_x     = -63.152530670166016
  native_bounding_box_min_y     = -21.387481689453125
  native_bounding_box_crs_value = "EPSG:4326"

  projection_policy = "FORCE_DECLARED"

  srs = "EPSG:4326"

  attribute {
    name       = "geometrie"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "org.locationtech.jts.geom.Geometry"
  }
  attribute {
    name       = "cleabs"
    min_occurs = 1
    max_occurs = 1
    nillable   = false
    binding    = "java.lang.String"
  }
  attribute {
    name       = "nature"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "usage_1"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "usage_2"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "construction_legere"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.Boolean"
  }
  attribute {
    name       = "etat_de_l_objet"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "date_creation"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.sql.Timestamp"
  }
  attribute {
    name       = "date_modification"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.sql.Timestamp"
  }
  attribute {
    name       = "date_d_apparition"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.sql.Date"
  }
  attribute {
    name       = "date_de_confirmation"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.sql.Date"
  }
  attribute {
    name       = "sources"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "identifiants_sources"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "precision_planimetrique"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "precision_altimetrique"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "nombre_de_logements"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.Integer"
  }
  attribute {
    name       = "nombre_d_etages"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.Integer"
  }
  attribute {
    name       = "materiaux_des_murs"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "materiaux_de_la_toiture"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "hauteur"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "altitude_minimale_sol"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "altitude_minimale_toit"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "altitude_maximale_toit"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "altitude_maximale_sol"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.math.BigDecimal"
  }
  attribute {
    name       = "origine_du_batiment"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
  attribute {
    name       = "appariement_fichiers_fonciers"
    min_occurs = 0
    max_occurs = 1
    nillable   = true
    binding    = "java.lang.String"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attribute` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--attribute))
- `name` (String)
- `native_name` (String)
- `projection_policy` (String)
- `srs` (String)
- `workspace_name` (String)

### Optional

- `abstract` (String)
- `datastore_name` (String)
- `enabled` (Boolean)
- `lat_lon_bounding_box_crs_class` (String)
- `lat_lon_bounding_box_crs_value` (String)
- `lat_lon_bounding_box_max_x` (Number)
- `lat_lon_bounding_box_max_y` (Number)
- `lat_lon_bounding_box_min_x` (Number)
- `lat_lon_bounding_box_min_y` (Number)
- `metadata` (Map of String)
- `native_bounding_box_crs_class` (String)
- `native_bounding_box_crs_value` (String)
- `native_bounding_box_max_x` (Number)
- `native_bounding_box_max_y` (Number)
- `native_bounding_box_min_x` (Number)
- `native_bounding_box_min_y` (Number)
- `native_crs_class` (String)
- `native_crs_value` (String)
- `title` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--attribute"></a>
### Nested Schema for `attribute`

Required:

- `binding` (String)
- `max_occurs` (Number)
- `min_occurs` (Number)
- `name` (String)
- `nillable` (Boolean)

